# myhttp

myhttp is a tool which make concurrent http requests and output the url of the request along with the MD5 hash of the response.

### Running

Clone the myhttp repo

`git clone https://github.com/nabilshafi/myhttp.git`

Go inside `myhttp` directory. Type following command to build the project

`go build myhttp.go`

Check the example for run the binary file

### Example

`./myhttp http://www.adjust.com http://google.com`

`./myhttp -parallel=3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com`

### Run Test

For executing the test

`go test -v`
