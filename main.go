package main

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"

	"os"

	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"
	_ "gocloud.dev/secrets/gcpkms"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cloudinary_name := os.Getenv("CLOUDINARY_NAME")
	cloudinary_api_key := os.Getenv("CLOUDINARY_API_KEY")
	cloudinary_api_secret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, _ := cloudinary.NewFromParams(cloudinary_name, cloudinary_api_key, cloudinary_api_secret)
	ctx := context.Background()
	images_path := "./images"

	filepath.Walk(images_path, func(file_path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if info.Name() != "images" {
			filename := info.Name()
			public_id := strings.TrimSuffix(filename, path.Ext(filename))
			resp, err := cld.Upload.Upload(ctx, file_path, uploader.UploadParams{PublicID: public_id,
				Transformation: "q_auto"})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp)
			}
		}

		return nil
	})

}
