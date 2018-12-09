Build image
`docker build -t tempio .`

Run image
`docker run --name="tempio-container" -d -v ~/tempio:/tmp --rm -it -p 8080:8080 tempio`
