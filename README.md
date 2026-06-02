# ASCII-Art-Web

## Description

ASCII-Art-Web is a web-based version of the ASCII-Art project written in Go. It allows users to enter text through a web interface and generate ASCII art using different banner styles.

Instead of running commands in the terminal, users interact with a graphical web page where they can submit text and select a banner template. The server processes the request and displays the generated ASCII art in the browser.

Supported banners:

* standard
* shadow
* thinkertoy

## Objectives

This project demonstrates:

* HTTP server creation in Go
* Handling GET and POST requests
* HTML templating
* Form processing
* ASCII art generation

## Project Structure

ascii-art-web/
│
├── main.go
├── web-handler.go
├── ascii-art.go
│
├── templates/
│   └── index.html
│
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
│
├── go.mod
└── README.md

## HTTP Endpoints

### GET /

Displays the main page containing:

* Text input field
* radio buttons to select banners
* Generate button

### POST /

Receives:

* User text
* Selected banner

## Installation

Clone the repository:

git clone <repository-url>
cd ascii-art-web

## Running the Application

go run .

### The application will be available at:

http://localhost:8080

Open your browser and visit the address above.

## Usage

1. Open the application in your browser.
2. Enter the text you want to convert.
3. Select a banner style:

   * standard
   * shadow
   * thinkertoy
4. Click **Generate**.
5. View the generated ASCII art on the page.

## Technologies Used

* Go
* HTML5
* Go Templates
* HTTP Package


