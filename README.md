# Taxi Fares

Program that gets taxi fares in Go.

## How to run:

- Run using `go run`
- Build a binary with `go build` and run it 
- Run using Docker

## Using `go run` command

1. Make sure you already install Go 1.21 <https://go.dev/dl> in your machine>
2. Navigate to the project directory: e.g. `cd taxi-fares`
3. Run command : `go mod download` to download all dependencies
4. Run command : `go run main.go` to run the program

## Build a binary with `go build` and run it

1. Make sure you already install Go 1.21 <https://go.dev/dl> in your machine>
2. Navigate to the project directory: e.g. `cd taxi-fares`
3. Run command : `go mod download` to download all dependencies
4. Run command : `go build -o {binary_output_name} .` e.g. `go build -o main .`
5. Execute your binary 
   1. If you are using mac/linux you can run command `./{binary-name}` e.g. `./main`
   2. If you are using windows you can run command `start {file_name.exe}` e.g. `start main.exe`

## Running with Docker

1. Make sure you already install Docker <https://docs.docker.com/get-docker> in your machine>
2. Navigate to the project directory: e.g. `cd taxi-fares`
3. Run command : `docker build -t taxi-fares .` to create docker image
4. Run command : `docker run -i -t -v {dir_on_your_machine}:/app/{log_dir} taxi-fares` to run the program e.g. `docker run -i -t -v ./log:/app/log taxi-fares`
   1. `{dir_on_your_machine}` : directory to store log file on your machine
   2. `{log_dir}` : log directory path, same with `file` on `config.yaml` but without filename e.g. in your `config.yaml` the value is `log/app.log` you need to input only `log` 

