package consumer_test

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/SinisterSup/messageQueuing/internal/app/consumer"
)

func TestCreateImageDirectory(t *testing.T) {
	dir := consumer.CreateImageDirectory()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("Expected the directory to be created, but it was not: %v", err)
	}

	err := os.RemoveAll(dir)
	if err != nil {
		t.Errorf("Failed to remove the created directory: %v", err)
	}
}

func TestDownloadAndCompressImages(t *testing.T) {
	// Replace these URLs with valid image URLs for testing
	imageURLs := []string{
		"https://images.everydayhealth.com/images/what-are-natural-skin-care-products-alt-1440x810.jpg",
		"https://www.shutterstock.com/image-vector/set-different-beauty-cosmetic-products-260nw-1942263496.jpg",
	}

	paths, err := consumer.DownloadAndCompressImages(imageURLs)
	if err != nil {
		t.Errorf("Expected the images to be downloaded and compressed, but got an error: %v", err)
	}
	imgPath := strings.Join(paths[:len(paths)-1], "\\")
	// check if the current path is the same as the path of imgPath
	if filepath.Dir(imgPath) != strings.Join(paths[:len(paths)-2], "\\") {
		t.Errorf("Expected the images to be downloaded and compressed, but got an error: %v", err)
	}
	
}
