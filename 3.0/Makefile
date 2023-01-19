CURRENT_DIR := $(dir $(MAKEFILE_PATH))
specs := /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/info.yaml /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/flatten_paths.yaml
flatten:
	swagger303 flatten /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/path.yaml --format=yaml -o /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/flatten_paths.yaml

version:
	@echo "##################################################################################################"
	swagger303 version
	@echo "##################################################################################################"
mixin: flatten version
	echo ${specs}
	swagger303 mixin ${specs} --format yaml -o /Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/swag_api.yaml
generate: mixin
	swagger303 generate server --target=./gen --spec=/Users/shantanu/SpectroCloud/go-swagger-demo/3.0/spec/swag_api.yaml --name=SwagApi
