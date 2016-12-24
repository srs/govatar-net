package main

import (
	"bytes"
	"crypto/md5"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/o1egl/govatar"
)

var (
	// Version set by build.
	Version = "x.x.x"
)

func parseGender(str string) (govatar.Gender, []int) {
	maxAssets := make([]int, 6)
	maxAssets[0] = len(govatar.Data.Background)

	if str == "f" {
		maxAssets[1] = len(govatar.Data.Female.Face)
		maxAssets[2] = len(govatar.Data.Female.Clothes)
		maxAssets[3] = len(govatar.Data.Female.Mouth)
		maxAssets[4] = len(govatar.Data.Female.Hair)
		maxAssets[5] = len(govatar.Data.Female.Eye)
		return govatar.FEMALE, maxAssets
	} else {
		maxAssets[1] = len(govatar.Data.Male.Face)
		maxAssets[2] = len(govatar.Data.Male.Clothes)
		maxAssets[3] = len(govatar.Data.Male.Mouth)
		maxAssets[4] = len(govatar.Data.Male.Hair)
		maxAssets[5] = len(govatar.Data.Male.Eye)
		return govatar.MALE, maxAssets
	}
}

func findPart(val byte, maxNum int) int {
	return int((val & 0x8f) % byte(maxNum))
}

func calcAssets(str string, maxAssets []int) []int {
	assets := make([]int, 6)
	hash := md5.Sum([]byte(str))

	assets[0] = findPart(hash[0], maxAssets[0])
	assets[1] = findPart(hash[1], maxAssets[1])
	assets[2] = findPart(hash[2], maxAssets[2])
	assets[3] = findPart(hash[3], maxAssets[3])
	assets[4] = findPart(hash[4], maxAssets[4])
	assets[5] = findPart(hash[5], maxAssets[5])

	return assets
}

func encodeImage(img image.Image, ext string) (*bytes.Buffer, string, error) {
	buffer := new(bytes.Buffer)

	if ext == "png" {
		e := png.Encode(buffer, img)
		return buffer, "image/png", e
	}

	e := jpeg.Encode(buffer, img, nil)
	return buffer, "image/jpeg", e
}

func writeImage(writer http.ResponseWriter, image image.Image, ext string) {
	buffer, contentType, _ := encodeImage(image, ext)

	writer.Header().Set("Content-Type", contentType)
	writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	writer.Write(buffer.Bytes())
}

func resizeImage(img image.Image, size string) image.Image {
	if size == "" {
		return img
	}

	width, _ := strconv.Atoi(size)
	if width == img.Bounds().Max.X {
		return img
	}

	return resize.Resize(uint(width), 0, img, resize.Lanczos3)
}

func serveAvatar(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	gender, maxAssets := parseGender(vars["gender"])
	assets := calcAssets(vars["hash"], maxAssets)

	result, err := govatar.GenerateFromAssets(gender, assets)
	if err != nil {
		panic(err)
	}

	img := resizeImage(result, request.URL.Query().Get("size"))
	writeImage(writer, img, vars["ext"])
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := mux.NewRouter()
	router.HandleFunc("/{gender}/{hash}.{ext}", serveAvatar).Methods("GET")

	addr := ":" + port
	log.Println("Govatar Net " + Version)
	log.Println("Serving avatars from " + addr + "...")
	log.Fatal(http.ListenAndServe(addr, router))
}
