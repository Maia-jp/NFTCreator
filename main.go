package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var target, output, name string
	var piece, threads int

	cliTitle()
	cliMenuWelcome()

	target, output = cliSelectFolder()
	fmt.Println(target, output)
	var tnm *NFT = newTargetNFTMeta(target, output)

	cliDisplayInfo(tnm)
	name, piece = cliCollectionName()
	threads = cliThread()

	cliReady()
	generateNFT(tnm, name, piece, threads)

	cliFinish(tnm)

	return

}
