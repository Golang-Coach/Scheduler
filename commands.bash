go get github.com/onsi/ginkgo/ginkgo

ginkgo -r -cover

overalls -project=Scheduler -debug -- -race -v

go tool cover -html=overalls.coverprofile

