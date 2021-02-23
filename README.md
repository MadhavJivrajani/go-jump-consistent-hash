# Jump Consistent Hashing

This is a Golang implementation of a *A Fast, Minimal Memory, Consistent Hash Algorithm* called *The Jump Consistent Algorithm*.
Ref: https://arxiv.org/pdf/1406.2294.pdf

## Example:
```go
import (
    "fmt"
    jch "github.com/MadhavJivrajani/go-jump-consistent-hash"
)

func main() {
    // a dummy stream of keys
    stream := []unit64{132415, 134131, 223, 6, 1, 76, 446, 8888}
    // number of buckets available
    var numBuckets int32 = 7
    for _, ele := range stream {
        // calculate the bucket number of the key
        // should be in range [0, 6)
        err := process(Hash(ele, numBuckets))
        if err != nil {
            panic(err)
        }
    }
}

// do some processing with the calculated bucket number
func process(bucket int64) error {
    fmt.Println("Bucket:", bucket)
    return nil
}
```
## Contributing
Please feel free to raise an issue if you find something that can be implemented in a better way or if something is just wrong. 
### TODO
- [ ] Provide conversion functions to convert keys to `uint64`