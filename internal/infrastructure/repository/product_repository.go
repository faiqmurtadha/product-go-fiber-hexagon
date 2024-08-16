package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"product-go-fiber-hexagon/internal/core/model"
	"product-go-fiber-hexagon/internal/core/port"
	"runtime/pprof"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) port.ProductRepository {
	return &ProductRepository{collection: collection}
}

func (repo *ProductRepository) profilingWrapper(operation string, f func() error) error {
	// Start CPU profiling
	cpuFile, err := os.Create(fmt.Sprintf("cpu_profiling_%s.prof", operation))
	if err != nil {
		log.Printf("Could not create CPU profile: %v", err)
	} else {
		pprof.StartCPUProfile(cpuFile)
		defer pprof.StopCPUProfile()
	}

	// Start timing
	start := time.Now()

	// Execute the operation
	err = f()

	// Record duration
	duration := time.Since(start)

	// Log duration
	log.Printf("MongoDB %s operation took %v", operation, duration)

	// Write memory profile
	memFile, err := os.Create(fmt.Sprintf("mem_profile_%s.prof", operation))
	if err != nil {
		log.Printf("Could not create memory profile: %v", err)
	} else {
		defer memFile.Close()
		if err := pprof.WriteHeapProfile(memFile); err != nil {
			log.Printf("Could not write memory profile: %v", err)
		}
	}

	return err
}

func (repo *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	var result *model.Product
	err := repo.profilingWrapper("createData", func() error {
		product.ID = primitive.NewObjectID()
		_, err := repo.collection.InsertOne(context.TODO(), product)
		if err != nil {
			return err
		}
		result = product
		return nil
	})

	return result, err
}

func (repo *ProductRepository) GetAll(page, limit int64, name string) (*model.PaginatedProduct, error) {
	var result *model.PaginatedProduct
	err := repo.profilingWrapper("getAllProduct", func() error {
		ctx := context.Background()

		skip := (page - 1) * limit

		filter := bson.M{}
		if name != "" {
			filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: "^" + strings.TrimSpace(name), Options: "i"}}
		}

		// Get total count
		totalCount, err := repo.collection.CountDocuments(ctx, filter)
		if err != nil {
			return err
		}

		// Get paginated products
		opts := options.Find().SetSkip(skip).SetLimit(limit)
		cursor, err := repo.collection.Find(ctx, filter, opts)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		var products []model.Product
		err = cursor.All(ctx, &products)

		result = &model.PaginatedProduct{
			Products:   products,
			TotalCount: totalCount,
			Page:       page,
			Limit:      limit,
		}

		return err
	})

	return result, err
}

func (repo *ProductRepository) FindById(id string) (*model.Product, error) {
	var product model.Product
	err := repo.profilingWrapper("getProductById", func() error {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		err = repo.collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&product)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) Update(id string, product *model.Product) (*model.Product, error) {
	var result *model.Product
	err := repo.profilingWrapper("updateProduct", func() error {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		filter := bson.M{"_id": objId}

		update := bson.M{"$set": product}

		_, err = repo.collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return err
		}

		product.ID = objId
		result = product
		return nil
	})

	return result, err
}

func (repo *ProductRepository) Delete(id string) error {
	return repo.profilingWrapper("deleteProduct", func() error {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		_, err = repo.collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
		return err
	})
}
