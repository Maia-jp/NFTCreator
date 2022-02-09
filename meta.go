// =========Description=========
// Handles all "meta" information - Loads the layers and some configuration

package main

import (
	"io/ioutil"
	"log"
	"math"

	WeightedRandomChoice "github.com/kontoulis/go-weighted-random-choice"
)

//NFT information
type NFT struct {
	wSize         int
	hSize         int
	targetPath    string
	resultingPath string
	nLayers       int
	nFiles        int
	layers        [][]string
	probchoice    []WeightedRandomChoice.WeightedRandomChoice
}

//NFT metadata to be writen in json file
type jsonNFT struct {
	Name         string            `json:"Name"`
	FilePath     string            `json:"FilePath"`
	CreationDate string            `json:"CreationDate"`
	Rarity       float32           `json:"Rarity"`
	Layers       map[string]string `json:"Layers"`
}

// Creates a new targetNFTMeta struct
//Parameters: Path to find all Layers, Path to save all information
//Return: Pointer to NFT struct
func newTargetNFTMeta(targetPath string, resultPath string) *NFT {
	var nLayers int
	var irow int = 0
	tnm := NFT{}
	tnm.nFiles = 0

	//Find the number of layers in the root dir
	root, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	nLayers = len(root)
	tnm.nLayers = nLayers
	tnm.layers = make([][]string, nLayers)

	//General information and weigthed choice obj
	tnm.targetPath = targetPath
	tnm.resultingPath = resultPath
	for i := 0; i < nLayers; i++ {
		tnm.probchoice = append(tnm.probchoice, WeightedRandomChoice.New())
	}

	//Apend image paths to each layer
	for _, f := range root {
		layerPath, err := ioutil.ReadDir(targetPath + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		var n int = 1
		if f.IsDir() {
			for _, l := range layerPath {
				if !l.IsDir() {
					var path string = targetPath + "/" + f.Name() + "/" + l.Name()
					tnm.layers[irow] = append(tnm.layers[irow], path)
					tnm.probchoice[irow].AddElement(path, generateProb(n, 0))
					n++
					tnm.nFiles++
				}
			}
		}
		irow++
	}

	tnm.hSize = -1
	tnm.wSize = -1

	return &tnm
}

// Generates the probability of each image to be choosen. Based on File order
//Parameters: File index, Enum of formulas (Linear=0, exp=1)
//Return: Probability of be choosen
func generateProb(n int, formula int) int {
	switch formula {
	//Linear
	case 0:
		return n * 2
	//Exp
	case 1:
		var r int
		r = int(math.Pow(2, math.Log(float64(n)))) * 3
		return r
	}
	return n
}
