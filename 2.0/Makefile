CURRENT_DIR := $(dir $(MAKEFILE_PATH))
specs := /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/info.yaml /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/flatten_paths.yaml
flatten:
	swagger flatten /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/path.yaml --format=yaml -o /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/flatten_paths.yaml

version:
	@echo "##################################################################################################"
	swagger version
	@echo "##################################################################################################"
mixin: flatten version
	echo ${specs}
	swagger mixin ${specs} --format yaml -o /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/swag_api.yaml
generate: mixin
	swagger generate server --target=./gen --spec=/Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/swag_api.yaml --name=SwagApi
