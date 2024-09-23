# Lignaj

## Introduction

Lignaj is a useful tool when working on movie dubbing.  
The tool will take a directory containing a list of files, one for each chapter.  
Each chapter has a list of characters with how many lines of text they are speaking.  

The output will be a csv file with the name of the character and the total number of lines across all chapters.

## Basic Usage

### Prerequisites

You will need to have [Go](https://go.dev/) installed on your computer first.  
The easiest way is to use [homebrew](https://brew.sh/) for that:
```bash
brew install golang
```

Also, you need to have [Make](https://www.gnu.org/software/make/) installed on your computer.  
Again, you can use [homebrew](https://brew.sh/) for that.
```bash
brew install make
```

### Installing lignaj

Then you can simply use the Makefile to install the utility:
```bash
make install
```

### Converting all files in directory

Once the utility is installed you can run it to convert all files in directory.  
The first argument will the directory containing the files.
```bash
lignaj -i . -o output.csv
```

The flag `-i` is used to indicate the input directory, which contains the files to process.  
The flag `-o` is used to indicate the output file.

The output file will have a similar structure:
```csv
Anv,Niver
Michael Scott,42
Dwight Schrute,13
Jim Halpert,5
```

## Contributing

Contributions are welcome!  
Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## Licenses

This program is under the terms of the BSD 3-Clause License.  
See [https://opensource.org/licenses/BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause).