// =========Description=========
// Join all images into a single one. Creates the metadata to be writen in json
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

//Aux Function to find an index of an element
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

//Generates the NFTs
//Parameters: NFT Struct, Number of images to be created, number of threads
//Return: NULL
func generateNFT(nft *NFT, name string, outSize int, WORKER int) {
	var wg sync.WaitGroup

	getOutSize(nft)
	var slices int = outSize / WORKER

	for i := 0; i < WORKER; i++ {
		wg.Add(1)
		ii := i
		go func() {
			defer wg.Done()
			generateNFTWORKER(nft, name, ii, ii*slices, ii*slices+slices)
		}()
	}

	//If there is any leftover
	if outSize%WORKER != 0 {
		wg.Add(1)
		ii := WORKER * slices
		go func() {
			defer wg.Done()
			generateNFTWORKER(nft, name, ii, ii, ii+outSize%WORKER+1)
		}()
	}

	wg.Wait()

}

//Worker Thread to generate NFTs
//Parameters: NFT Struct, routine id, start, finish
//Return: NULL
func generateNFTWORKER(nft *NFT, name string, id int, start int, finish int) {
	for j := start; j < finish; j++ {
		newNFT(nft, fmt.Sprintf("%d", j), name)

	}
}

//Get the base layer image size
//Parameters: NFT Struct
//Return: NULL
func getOutSize(nft *NFT) {
	if nft.wSize > 0 || nft.hSize > 0 {
		return
	}
	reader, err := os.Open(nft.layers[0][0])
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(reader)
	if err != nil {
		fmt.Print("Err decode:")
		log.Fatal(err)
	}

	bounds := m.Bounds()
	nft.wSize = bounds.Dx()
	nft.hSize = bounds.Dy()

	reader.Close()
}

//Create and save an new nft
//Parameters: NFT Struct, Name of NFT, Post name of nft
//Return: NULL
func newNFT(nft *NFT, posfix string, sufix string) {
	//Metadata Info
	metadata := jsonNFT{}
	metadata.CreationDate = time.Now().String()
	metadata.Rarity = 1
	metadata.Layers = make(map[string]string)

	//Image info
	upLeft := image.Point{0, 0}
	lowRight := image.Point{nft.wSize, nft.hSize}

	//Generates the new image
	outImg := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	//Overlay layers
	for i := 0; i < nft.nLayers; i++ {
		//Get Layer
		layerName := nft.probchoice[i].GetRandomChoice()

		//Calculates Rarity
		wKey := indexOf(layerName, nft.probchoice[i].Elements)
		metadata.Rarity = metadata.Rarity * float32(nft.probchoice[i].Weights[wKey]) / float32(nft.probchoice[i].TotalWeight)
		layerKey := fmt.Sprintf("Layer #%d", i)
		metadata.Layers[layerKey] = layerName

		imgFile, err := os.Open(layerName)
		if err != nil {
			fmt.Println(err)
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
		}

		draw.Draw(outImg, img.Bounds(), img, image.Point{0, 0}, draw.Over)

		imgFile.Close()

	}

	//Generate name
	var name string = fmt.Sprintf("%s #%s", sufix, posfix)
	out, err := os.Create(nft.resultingPath + "/" + name + ".jpeg")
	if err != nil {
		fmt.Println(err)
	}

	//Generates Json path
	metadata.Name = name
	metadata.FilePath = nft.resultingPath + "/Metadata/" + name + ".json"
	_ = os.Mkdir(nft.resultingPath+"/Metadata", 0644)

	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(out, outImg, &opt)
	out.Close()

	//Save Metadata
	file, _ := json.MarshalIndent(metadata, "", " ")
	_ = ioutil.WriteFile(metadata.FilePath, file, 0644)
}
