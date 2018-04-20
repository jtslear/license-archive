# license_archive
## Purpose
* Help me archive older sets of license quickly based on criterion we observe

## How it Works
* Queries the [Replicated API](https://replicated-vendor-api.readme.io/v1.0/reference)
  for all our licenses in the application
* Using my super duper complicated if statement, rolls through and determines
  which should be archived
* After I complete this, it'll then archive them (TODO)

## Get Started
* `brew install go`
* Setup your go environment, utilize golang docs for this
* `go get github.com/stretchr/testify/assert`
* in your go `src` path: `git clone git@jtslear/license-archive`
* `go test`

## Execution
* I'm not building a binary for this, so: `go run main.go`

## Feature Requests
* Create a legit cli to provide various options
  * dry-run capability
  * verbosity
* Better manage dependencies
* Better the method of filtering down the map of data
* Add an undo capability, just in case I screw up badly
