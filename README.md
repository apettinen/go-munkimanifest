# go-munkimanifest

Create a [munki](https://github.com/munki/munki) manifest based on template. Usable if you want to e.g. automate your manifest creation.

## Usage

To use this program, you need to build it via `go build munkimanifest.go`. 
To build for another architecture, set `GOOS` and `GOARCH` appropriately. For example, to build for Linux, run `GOOS=linux GOARCH=amd64 go build -o munkimanifest munkimanifest.go`.


To actual usage instructions, see `munkimanifest -h`


## License:

Apache 2.0 Licensed, for more info, see LICENSE
