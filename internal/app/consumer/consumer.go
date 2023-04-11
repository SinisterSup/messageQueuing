package consumer

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nfnt/resize"

	"github.com/SinisterSup/messageQueuing/internal/db"
	"github.com/SinisterSup/messageQueuing/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DownloadCompressAndStoreImages(client *mongo.Client, productID primitive.ObjectID) error {
	productCollection := db.GetCollection(client, "UserProduct-Messaging", "Products")

	var product models.Product
	err := productCollection.FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)

	if err != nil {
		log.Printf("Failed to find the product: %v", err)
		return err
	}

	// downloadAndCompressImages function takes in []string of image urls and 
	// returns a []string of local paths of compressed images.
	imgsPath, err := DownloadAndCompressImages(product.Images)
	if err != nil {
		log.Printf("Failed to download and compress images: %v", err)
		return err
	}

	// Update the product with compressed images in the database
	_, err = productCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": productID},
		bson.D{
			bson.E{Key: "$set", Value: bson.D{{Key: "compressed_product_images", Value: imgsPath}}},
		},
	)

	if err != nil {
		log.Printf("Failed to update the product with compressed images: %v", err)
		return err
	}

	fmt.Printf("Updated product %s with compressed images\n", productID)
	return nil
}

func DownloadAndCompressImages(imageURLs []string) ([]string, error) {
	imgDir := CreateImageDirectory()
  	absPath, err := filepath.Abs(imgDir)
  	if err != nil {
		log.Printf("Failed to get the absolute path of the image directory: %v", err)
		return nil, err
  	}
	imgDirPath := strings.Split(filepath.ToSlash(absPath), "/")

	// var imageFolderPaths []string
	for i, url := range imageURLs {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to download the image from the URL: %v", err)
			return nil, err
		}
		defer resp.Body.Close()

		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Printf("Failed to decode the image: %v", err)
			return nil, err
		}

		compressedImg := resize.Resize(0, 0, img, resize.Lanczos3)

		imgPath := filepath.Join(imgDir, fmt.Sprintf("img-%d.jpg", i))
		out, err := os.Create(imgPath)
		if err != nil {
			log.Printf("Failed to create the compressed image file: %v", err)
			return nil, err
		}
		defer out.Close()

		err = jpeg.Encode(out, compressedImg, &jpeg.Options{Quality: 75})
		if err != nil {
			log.Printf("Failed to encode the compressed image: %v", err)
			return nil, err
		}
	}

	return imgDirPath, nil
}

func CreateImageDirectory() string {
	dir := filepath.Join("imgs", strconv.FormatInt(time.Now().Unix(), 10))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create image directory: %v", err)
	}
	return dir
}