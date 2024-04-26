# Go Templ/Htmx Base Project

My base project for using Go, Templ, and HTMX

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

- Install and setup [air](https://github.com/cosmtrek/air)
- Install and setup [nodejs](https://nodejs.org/en/download)
- Install templ cli ```go install github.com/a-h/templ/cmd/templ@latest```
- run ```npm install```
- run ```go mod download```
- run ```air```

## Usage

This project is a base project using the following technologies:

- [GO](https://go.dev/)
- [Templ](https://templ.guide/)
- [HTMX](https://htmx.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [DaisyUI](https://daisyui.com/)
- [air](https://github.com/cosmtrek/air) - Hot Reloading

After running the project you will have a logging middleware and a test middle ware that was added to push a static user into context. This users email is then displayed on the ```localhost:8080/user``` endpoint as an example.

## Notes

The cmd to run the project in the .air.toml is as follows: ```"npx tailwindcss -o static/styles.css --minify && templ generate && go build -o ./tmp/main ./main.go"```. This will do the following when running the project:

- Generate your tailwindcss minified css
- Generate your templ go classes
- Build and run your project
