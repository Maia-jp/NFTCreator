// =========Description=========
// IO used in the CLI
package main

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ttacon/chalk"
)

func cliTitle() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(` /$$   /$$ /$$$$$$$$ /$$$$$$$$        /$$$$$$                                                     /$$                        `)
	fmt.Println(`| $$$ | $$| $$_____/|__  $$__/       /$$__  $$                                                   | $$                        `)
	fmt.Println(`| $$$$| $$| $$         | $$         | $$  \__/  /$$$$$$  /$$$$$$$   /$$$$$$   /$$$$$$  /$$$$$$  /$$$$$$    /$$$$$$   /$$$$$$ `)
	fmt.Println(`| $$ $$ $$| $$$$$      | $$         | $$ /$$$$ /$$__  $$| $$__  $$ /$$__  $$ /$$__  $$|____  $$|_  $$_/   /$$__  $$ /$$__  $$`)
	fmt.Println(`| $$  $$$$| $$__/      | $$         | $$|_  $$| $$$$$$$$| $$  \ $$| $$$$$$$$| $$  \__/ /$$$$$$$  | $$    | $$  \ $$| $$  \__/`)
	fmt.Println(`| $$\  $$$| $$         | $$         | $$  \ $$| $$_____/| $$  | $$| $$_____/| $$      /$$__  $$  | $$ /$$| $$  | $$| $$      `)
	fmt.Println(`| $$ \  $$| $$         | $$         |  $$$$$$/|  $$$$$$$| $$  | $$|  $$$$$$$| $$     |  $$$$$$$  |  $$$$/|  $$$$$$/| $$      `)
	fmt.Println(`|__/  \__/|__/         |__/          \______/  \_______/|__/  |__/ \_______/|__/      \_______/   \___/   \______/ |__/      `)
	fmt.Println(chalk.Underline.TextStyle("By JP Maia"))
}

func cliMenuWelcome() {
	fmt.Println("Welcome to the NFT Generator. This is a tool to create layer-Based NFTs.")
	fmt.Println("Rarity levels and metadata will be created too. Custom attributes are not avaliable at the moment")
	fmt.Println("Read the fast documentation before using it")
	fmt.Println("-------------------------------------------")
	fmt.Println("Press enter to start...")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")

}

func cliSelectFolder() (string, string) {
	//Source folder
	pathTarget := ""
	for true {
		prompt := &survey.Input{
			Message: "The fist step is to write (or paste) the path of the layers folder (contaning the layers as sub-folders):",
		}
		survey.AskOne(prompt, &pathTarget)

		//Default path
		if len(pathTarget) < 1 {
			pathTarget = "../Example/target"
			break
		}

		fmt.Print("\033[H\033[2J")
		fmt.Println("You choose the following path: ", chalk.Cyan, pathTarget, chalk.Reset)

		ok := false
		prompt2 := &survey.Confirm{
			Message: "This is the right path ?",
		}
		survey.AskOne(prompt2, &ok)

		if ok {
			fmt.Println("Nice! Now wee need output folder")
			break
		} else {
			fmt.Println("Ok!")
		}
	}

	//outputfolder
	pathResult := ""
	for true {
		prompt := &survey.Input{
			Message: "Write (or paste) the output path :",
		}
		survey.AskOne(prompt, &pathResult)
		//Default path
		if len(pathResult) < 1 {
			pathResult = "../Example/result"
			break
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println("You choose the following path: ", chalk.Cyan, pathTarget, chalk.Reset)
		ok := false
		prompt2 := &survey.Confirm{
			Message: "This is the right path ?",
		}
		survey.AskOne(prompt2, &ok)

		if ok {
			fmt.Println("Nice! Reading folder now...")
			break
		} else {
			fmt.Println("Ok!")
		}
	}

	return pathTarget, pathResult
}

func cliDisplayInfo(nft *NFT) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Woah, we read %d files in %d diferent layers!\n", nft.nFiles, nft.nLayers)
}

func cliCollectionName() (string, int) {
	name := ""
	prompt := &survey.Input{
		Message: "Please write the collection name:",
	}
	survey.AskOne(prompt, &name)

	piecesSTR := ""
	prompt2 := &survey.Input{
		Message: "How many pieces you wish to create:",
	}
	survey.AskOne(prompt2, &piecesSTR)
	pices, err := strconv.Atoi(piecesSTR)

	if err != nil {
		pices = -1
	}

	return name, pices
}

func cliThread() int {
	fmt.Print("\033[H\033[2J")
	fmt.Println("We are almost done - Now it is time to choose how much power will be used")
	fmt.Println("A higher number will be fast. However your PC needs to handle it")
	fmt.Println("Do not choose a number higher than the number of cores in your PC")
	threadsStr := ""
	prompt := &survey.Select{
		Message: "How many Threads to use (press enter to skip):",
		Options: []string{"4", "6", "8", "12"},
		Default: "4",
	}
	survey.AskOne(prompt, &threadsStr)

	threads, err := strconv.Atoi(threadsStr)
	if err != nil {
		threads = 0
	}
	return threads
}

func cliReady() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Everything is ready, generating your NFTs ! This can take a while...")
}

func cliFinish(nft *NFT) {
	fmt.Print("\033[H\033[2J")
	fmt.Println(chalk.Green, "DONE ! No erros found", chalk.Reset)
	fmt.Println("Congradulations, your NFT collection has been created.")
	fmt.Printf("\nCheck your collection at %s", nft.resultingPath)
	fmt.Println("Closing the CLItool now...")
}
