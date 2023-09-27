# mailman

## named after karl malone, aka "the mailman", because he always delivered. 
*just a file sender with ftp.*

change the empty variables below to your ftp credentials and your local directory

	tgHostname := ""
	tgUsername := ""
	tgPassword := ""

    localDir := ""

clone the repo and run:

    go env -w GO111MODULE=off
    go get github.com/jlaffaye/ftp
    go run main.go

that's it