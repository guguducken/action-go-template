default: build

# You must modify it to yourself information
AUTHOR=guguducken
REPOSITORY=action-go-template
DESCRIPTION=A template repository for creating action written in go language.
ICON=arrow-up
COLOR=blue


.PHONY: build
build:
	@CGO_ENABLED=0 go build -o action .


.PHONY: clean
clean:
	@rm -rf action action.yaml


ACRION_GOLDFLAGS=-ldflags="-X 'main.Name=$(REPOSITORY)' -X 'main.Author=$(AUTHOR)' -X 'main.Description=$(DESCRIPTION)' -X 'main.Icon=$(ICON)' -X 'main.Color=$(COLOR)'"
.PHONY: action
action:
	@CGO_ENABLED=0 go run $(ACRION_GOLDFLAGS) ./optools/action