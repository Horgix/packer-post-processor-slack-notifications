build:: fmt
	go build -o packer-post-processor-slack-notifications

fmt::
	gofmt -w main.go ./slack-notifications

test::
	export PACKER_LOG=true
	packer build packer.json
