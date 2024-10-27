# ASCII Art Web Export File

## Overview
Ascii-art-web is a web application that allows users to generate ASCII art using different banners. It features a graphical user interface for inputting text and selecting a banner style.

## Installation
To use  ASCII Art Web Export File, you'll need Go installed on your system. You can install it via the official Go website.

Clone the repository:
```
git clone https://learn.zone01kisumu.ke/git/svictor/ascii-art-web-export-file
## Usage
Navigate to the project directory and run the program with the desired below commandline argument:
```
go run main.go
```
Open your browser and navigate to http://localhost:8082


## Project structure
```
* main.go: Contains the main program logic.
* banners/: Directory containing ASCII banner templates.
* templates/: contains the index.html and statis.css
* ascii/: containg go files that are used to generate ascii art
``` 

 ## project Implementation
```
   * The server is implemented in Go.
   * ASCII art generation logic is in ascii/
   * HTTP handlers are in main.go/.
   * HTML templates are in templates/.
```

  ## HTTP Status Codes
```
    * 200 OK: Successful response.
    * 404 Not Found: Resource not found.
    *  400 Bad Request: Incorrect request.
    * 500 Internal Server Error: Unhandled server error.
```

## Contributing
Contributions are welcome! If you have suggestions for improvements, please open an issue or create a pull request on the GitHub repository.


## Contributers
* [svictor](https://learn.zone01kisumu.ke/git/svictor)
* [bobaigwa](https://learn.zone01kisumu.ke/git/bobaigwa)
