package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	res "github.com/alifudin-a/go-ocr/http"
	"github.com/bregydoc/gtranslate"
	"github.com/labstack/echo/v4"
	"github.com/otiai10/gosseract"
	"golang.org/x/text/language"
)

// Convert image character to character
func Convert(c echo.Context) (err error) {
	var resp res.Response

	img, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Open file
	src, err := img.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// File destination
	dir := filepath.Join("./image", filepath.Base(img.Filename))
	dst, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//Tesseract-OCR
	client := gosseract.NewClient()
	defer client.Close()

	errResp := map[string]interface{}{
		"message": "An error occured! Make sure you have internet connection and a clear image file!",
	}

	imgPath := dir

	client.SetImage(imgPath)
	text, err := client.Text()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	textTranslate := &text
	translated, err := gtranslate.Translate(*textTranslate, language.English, language.Indonesian)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, errResp)
	}

	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("File %s uploaded successfully!", img.Filename)
	resp.Data = map[string]interface{}{
		"file":       img.Filename,
		"text":       text,
		"translated": translated,
	}

	return c.JSON(http.StatusOK, resp)
}
