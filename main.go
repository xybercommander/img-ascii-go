package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math"
	"os"
)

func main() {

	file, err := os.Open("bg.jpg")
	if err != nil {
		fmt.Println("Error opening image")
	}
	defer file.Close()

	// fmt.Println(file.Name())
	
	// Decoding the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image", err)
	}		
	
	
	// Converting into pixel array
	var pixels [][]Pixel
	height := img.Bounds().Max.X
	width := img.Bounds().Max.Y

	for i := 0; i < height; i++ {
		var row []Pixel
		for j := 0; j < width; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			var pixelVal Pixel = Pixel{
				int(r / 257),
				int(g / 257),
				int(b / 257),
				int(a / 257),
			}
			row = append(row, pixelVal)
		}
		pixels = append(pixels, row)
	}



	// Creating a new file
    newFile, err := os.Create("data.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer newFile.Close() // Ensure file is closed later


	verticalImageCharacters := [][]string{}

	for i := 0; i < height; i++ {
		rowASCIIdata := ""
		for j := 0; j < width; j++ {
			// fmt.Print(pixelToChar(pixels[i][j]))
			rowASCIIdata += pixelToChar(pixels[i][j])
		}
		rowASCIIdata += "\n"
		
		strArr := []string{}
		for j := 0; j < len(rowASCIIdata); j++ {
			strArr = append(strArr, string(rowASCIIdata[j]))
		}
		verticalImageCharacters = append(verticalImageCharacters, strArr)
		// _, err = newFile.WriteString(rowASCIIdata)
	}	

	newMatrix := transpose(verticalImageCharacters)

	
	for i := 0; i < len(newMatrix); i += 2 {
		r := ""
		for j := 0; j < len(newMatrix[i]); j++ {
			r += newMatrix[i][j]
		}
		r += "\n"
		_, err = newFile.WriteString(r)
	}


}


type Pixel struct {
	R int
	G int
	B int
	A int
}


// Calculate the luminance of the pixel
func luminance(p Pixel) float64 {
    return 0.299*float64(p.R) + 0.587*float64(p.G) + 0.114*float64(p.B)
}

// Map the luminance to a character
func charFromLuminance(lum float64) string {
    // Define a comprehensive set of characters from less vibrant to more vibrant
    chars := " .:-=+*%@#ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    index := int(math.Round(lum / 256.0 * float64(len(chars)-1)))
    return string(chars[index])
}

// Convert pixel to character based on its lightness
func pixelToChar(p Pixel) string {
    lum := luminance(p)
    return charFromLuminance(lum)
}

func transpose(matrix [][]string) [][]string {
	if len(matrix) == 0 {
		return [][]string{}
	}

	// Determine the size of the transposed matrix
	rows, cols := len(matrix), len(matrix[0])
	transposed := make([][]string, cols)
	for i := range transposed {
		transposed[i] = make([]string, rows)
	}

	// Fill the transposed matrix
	for i := range matrix {
		for j := range matrix[i] {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}