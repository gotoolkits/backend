## backend

A collection of modules to make backend programming easier.

### golang.org/x dep install
```
mkdir $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
```

### Catalog
* [config](https://github.com/gotoolkits/backend#config-module)
* [connector](https://github.com/gotoolkits/backend#connector-module)
* [handler](https://github.com/gotoolkits/backend#handler-module)
* [logs](https://github.com/gotoolkits/backend#logs-module)
* [server](https://github.com/gotoolkits/backend#server-module)
* [storage](https://github.com/gotoolkits/backend#storage-module)
* [utils](https://github.com/gotoolkits/backend#utils-module)

### Modules

#### config module
configuration file parser, supporting formats:
* json
* ini (characters/lines followed by ';' will be considered as comments)

```go
package main
import (
	"github.com/gotoolkits/backend/config"
	"fmt"
)

fun main() {
	// json config parser
	rc, err := rrconfig.LoadJsonConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}
	v, err := rc.GetStringSlice("files.ufile")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	// ini config parser
	ic, err := rrconfig.LoadIniConfigFromFile("test.ini")
	if err != nil {
		panic(err)
	}
	// get value by key
	s, err := ic.Get("test.a")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	// dump
	fmt.Println(ic.Dump())
	
}
```

#### connector module
clients for third-party service, supporting services:
* redis
* zookeeper

```go
package main

import (
	"fmt"
	"github.com/gotoolkits/backend/connector/redis"
	"github.com/gotoolkits/backend/connector/zookeeper"
)

func main() {
	// redis connector
	err, rc := rrredis.GetRedisClient("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	exist, err := rc.HMExists("TEST:KEY", "fool")
	if err != nil {
		panic(err)
	}
	fmt.Println(exist)

	// zk connector
	err, zkc := rrzk.GetZkClient("192.168.150.74:2181,192.168.150.75:2181,192.168.150.132:2181")
	if err != nil {
		panic(err)
	}
	fmt.Println(zkc)
}
```

### handler module
tcp handler register

```go
package main

import (
	"github.com/gotoolkits/backend/handler"
	"fmt"
)

func echo(c interface{}, msg interface{}) {
	fmt.Println("test")
}

func main() {
	_, hr := rrhandler.CreateHandlerRegister()
	hr.Add("rrfp.ExampleEchoRequest", rrhandler.Handler(echo), 0*time.Second)
}
```

#### logs module
loggers, supporting list:
* console
* file
* elasticsearch
* jianliao
* websocket
* slack
* smtp


#### server module
tcp server

```go
package main
import (
	"github.com/gotoolkits/backend/server"
)
func main() {
	err, s := rrserver.CreateTCPServer("0.0.0.0", 8003)
	if err != nil {
	    panic(err)
	}
	s.Start()
}
```

#### storage module
storage sdks, supporting storage:
* LocalDisk
* UFile

```go
package main
import (
	"github.com/gotoolkits/backend/storage"
)

func main() {
	// ufile
	se := rrstorage.CreateUfileStorage("publickey",
		"privatekey",
		"bucketname",
		2)

	// download file for ufile storage
	_, err = se.Fetch("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// local disk
	ls := rrstorage.CreateLocalDiskStorage("/data/files/")
	if err := ls.Save([]bytes("hehe"), "test.txt"); err != nil {
		fmt.Println(err)
	}
}
```

#### utils module
A collection of tools, suporting list:
* uuid
* pprof

```go
package main

import (
	"fmt"
	"github.com/gotoolkits/backend/utils"
)

func main() {
	// uuid
	uuid := rrutils.NewV4().String()
	fmt.Println(uuid)

	// pprof
	rrutils.StartProfiling()
}
```
