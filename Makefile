SETUP : install1 install2 install3

install1 :
	go install golang.org/x/tools/cmd/goimports@latest

install2 :
	go install honnef.co/go/tools/cmd/staticcheck@latest

install3 :
	go install github.com/kisielk/errcheck@latest

PREP : importscheck verify errcheck
	
importscheck :
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/TakaTaka1/linebot_go"
verify :
	@staticcheck ./... && go vet ./...
errcheck :
	@errcheck ./...
