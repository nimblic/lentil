# lentil - golang beanstalkd client library

_This is the lentil library. For the the command line client see [github.com/nutrun/lentil/tree/master/lentil](https://github.com/nutrun/lentil/tree/master/lentil)_

## INSTALL
```bash
go get github.com/nutrun/lentil
```

Or, in $GOPATH/src:

```bash
git clone git://github.com/nutrun/lentil.git && cd lentil && go install
```

After installing, use godoc to view documentation. For example:

```bash
godoc -http=:6060
```

And open http://0.0.0.0:6060/pkg/github.com/nutrun/lentil/ in a browser.

## USAGE

### Example producer:
```go
package main

import(
	"github.com/nutrun/lentil"
	"log"
)

func main() {
	conn, e := lentil.Dial("0.0.0.0:11300")
	if e != nil {
		log.Fatal(e)
	}
	jobId, e := conn.Put(0, 0, 60, []byte("hello"))
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("JOB ID: %d\n", jobId)
}
```

#### Connection and Call retries

By default, Dial and any functions called on the connection will not retry.
You can specify a retry period for both the Dial and other calls using arguments
to the Dial() call.

For instace, to retry the dial for 10 minutes and 
retry any call (Put, Delete, Reserve) for 1 minute before returning
an error, use:

```
func main() {
	conn, e := lentil.Dial("0.0.0.0:11300", "10m", "1m")
	if e != nil {
		log.Fatal(e)
	}
	jobId, e := conn.Put(0, 0, 60, []byte("hello"))
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("JOB ID: %d\n", jobId)
}
```
Both Dial and Put() will retry every second for the given duration.


### Example consumer:
```go
package main

import(
	"github.com/nutrun/lentil"
	"log"
)

func main() {
	conn, e := lentil.Dial("0.0.0.0:11300")
	if e != nil {
		log.Fatal(e)
	}
	job, e := conn.Reserve()
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("JOB ID: %d, JOB BODY: %s", job.Id, job.Body)
	e = conn.Delete(job.Id)
	if e != nil {
		log.Fatal(e)
	}
}
```

## LINKS
* beanstalkd: http://kr.github.com/beanstalkd/
* beanstalkd source: https://github.com/kr/beanstalkd/
* beanstalkd protocol: https://github.com/kr/beanstalkd/blob/master/doc/protocol.txt
* kr's go beanstalk client: https://github.com/kr/beanstalk

## LICENCE
See [LICENSE](https://github.com/nutrun/lentil/blob/master/LICENSE),
it's the same as [beanstalkd's license](https://github.com/kr/beanstalkd/blob/master/LICENSE)
