
# NFT Creator

No one can deny the impact of NFTs in this new decade. This technology is changing many fields and will change even more. As humans, we are standing at the edge of a new world where digital and reality tangle themselves into our pockets and minds. 

This project is a CLI tool to create digital 2D artworks. The main method used is "layering". This technique merges image layers (characteristics) to form a new image. This CLI tool provides a fast and easy way to combine layers to create a collection with different rarity levels.


## Tech

GO is the main language in this project. It was chosen due to its performance. It is statically compiled and has amazing support for multi-threading (goroutines). 
The syntax is easy which allows other people to modify the code.


## Layers

This image is a good example of "layering". Wich a bunch of assets is merged to create an image. In this case, we have Hats and glasses for example.

![Examples](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQJtJmotqFWeRSz9MTDFPpoQKyn39C6qdPDtQ&usqp=CAU)

Our CLI tool does the difficult job and merges a collection of images into a  collection.  
## Build and run

The first step is to download and install GO (1.17 or above). 

After that, we have to install some dependencies. 

```bash
  cd NFTCreator
  go get -u ./...
```
Finally, we can compile to our OS using the following command 

```bash
  go build .
```
This will generate an executable file.
## Demo

![image](https://user-images.githubusercontent.com/57457490/153283560-f79eefd7-c166-4131-9234-6c3eb2c1bfd4.png)

First, we need to create a folder where all subfolders contain images. The subfolders will be our attributes. It is important to organize the folder name in order for the first subfolder to be the first layer. 
We should pay attention to making all images with the same size and transparent background. 

![image](https://user-images.githubusercontent.com/57457490/153283368-90d1646f-ad98-4f59-ad46-aa82c5e2281f.png)

The rarity of each attribute is defined by the following function:  
`2^(log(j)) * 3`
Where `j` is the index of the file (the first file in and subfolder is rarer than the last one)

After the build, you can run the executable file and follow the given instruction. You will be asked about some metadata information as well as some performance tuning.
![image](https://user-images.githubusercontent.com/57457490/153283866-c0acb73e-2c12-417f-abb5-e0c1a012d3cd.png)


All pieces will be in a new folder. Inside this folder, there is the /metadata subfolder. Where is possible to see all information related to an image. It will be a .json like this:

```json
  {
 "Name": "0 #Test",
 "FilePath": "./Example/result/Metadata/0 #Test.json",
 "CreationDate": "2022-02-09 17:17:38.0307459 -0300 -03 m=+51.315811501",
 "Rarity": 0.00029629632,
 "Layers": {
  "Layer #0": "./Example/target/0_Head/Frame 1.png",
  "Layer #1": "./Example/target/1_Eyes/Frame 1.png",
  "Layer #2": "./Example/target/2_Mouth/Frame 1.png"
 }
}
```

Rarity is calculated by the multiplication of the probability of each layer appearing.

In the folder ./Example there is a couple of images that can be used in testing.

## Contributions
Contributions are always welcome! See the roadmap to have an idea of the next steps. After that feel free to open issues and fork


## Roadmap

- Better tests

- Custom attributes

- User adjusted rarity

- Solidity ready contract/metadata


