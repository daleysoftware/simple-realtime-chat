# As an exercise in Go dependency management, a self-contained Go environment is used here. We set
# GOPATH ourselves, install our dependency manager gvt, and use a little path hack to facilitate
# the build. Our dependencies are vendored and versioned via vendor/manifest. For the user of this
# repository, it should be as easy as `make`.

GOPATH="$(shell pwd)/.go"

all: setup run

run:
	cd "${GOPATH}/src/github.com/mpillar/simple-realtime-chat" && export GOPATH="${GOPATH}" && go run color.go

setup:
	mkdir -p "${GOPATH}"/src/github.com/mpillar
	cd "${GOPATH}/src/github.com/mpillar" && ln -s "$(shell pwd)" simple-realtime-chat
	export GOPATH="${GOPATH}" && go get -u github.com/FiloSottile/gvt
	export GOPATH="${GOPATH}" && "${GOPATH}"/bin/gvt update -all

clean:
	rm -rf "${GOPATH}"
	rm -rf vendor/github.com vendor/golang.org vendor/gopkg.in
