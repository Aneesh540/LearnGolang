## Golang all notes

#### Declare variables
```golang
s := "aneesh jain"
var a float64 = 3.45
```

#### Taking I/O
```golang

```

#### File handling Reading
```golang
content := "text content writtern to a file"
file, err := os.Create("./newfile.txt")

length, err := io.WriteString(file, content)
defer file.Close()
```

```golang
databyte, err := ioutil.Readfile(filepath)
return string(databyte)
```

#### Memory Management 

make : memory is allocated & initialized. Non zero storage
new : memory allocated but not initialized. Zeroed storage

#### GoRoutines
Threads managed by OS; 1MB size
GoRoutines managed by Go runtime: 2KB

#### Web request
```golang

import "net/http"

url := "http://ajs.com"
response, err := http.Get(url)
fmt.Printf("response type %T", response) // *http.Response

dataInBytes, err := ioutil.ReadAll(response.Body)
string(dataInBytes)

defer response.Body.Close() // caller responsible for closing conn.

```

