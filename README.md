# temap
Fast, Concurrency-safe, Timed Maps in Go

### What's a "timed map" ?
It's a simple `map` object which stores values for a specific amount of time, in memory.
pretty much like an in-memory cache store.

## Installation

```bash
go get -u github.com/firasdarwish/temap
```


## Usage

#### Initialising

```go
package main

import (
    temap "github.com/firasdarwish/temap"
    "time"
)

func main()  {
 // clean expired KVs every X (time.Duration)
 cleaningInterval := time.Minute * 2
 
 timedMap := temap.New(cleaningInterval)
}
```


#### Setting a temporary value
```go
    TTL := time.Second * 5
    expiresAt := time.Now().Add(TTL)

    timedMap.SetTemporary("age", 33, expiresAt)
```


#### Setting a permanent value
```go
    timedMap.SetPermanent("name", "Robert Langdon")
```


#### Retrieving a value by key
```go
    // `expiry` is a Unix timestamp in Nanoseconds,
    // if a value is set to be permanent then `expiry`=0.
    // if the value doesn't exists then `ok`=false & `expiry`=-1
    value, expiry, ok := timedMap.Get("age")
    if !ok {
        fmt.Println("value does not exists")    
    }else{
        fmt.Println(value)
    }
```


#### Remove a value by key
```go
    timedMap.Remove("name")
```


#### Remove all values
```go
    timedMap.RemoveAll()
```

#### Iterating over the map
```go
    // map[string]*element
    m := timedMap.ToMap()
    
    // timed map current elements count
    // mapSize := len(m)

    for key, element := range m {
        fmt.Println("KEY: "+key)
        fmt.Println("VALUE: "+element.Value)
        fmt.Println("EXPIRES AT: "+element.ExpiresAt)
    }

    // you can also marshal/unmarshal the timed map
    // b, err := json.Marshal(m)
```


#### Making a value; permanent
```go
    ok := timedMap.MakePermanent("age")
    if !ok {
        fmt.Println("value not found")
    }

    // OR
    age,_,ok := timedMap.Get("age")
    if ok {
        timedMap.SetPermanent("age", age)
    }
```

#### Setting a new expiration date
```go
    // can be used for both already temporary values & permanent values.
    newExpiry := time.Now().Add(time.Minute*10)
    ok := timedMap.SetExpiry("name", newExpiry)
    if !ok {
        fmt.Println("value not found")
}
```

### The Cleaner
By default, the cleaner starts working automatically
when initialising a new timed map,
and it will be triggered every X unit of time (time.Duration).

The cleaning operation is non-blocking for it is running on a separate goroutine.


#### Stopping the cleaner
```go
    timedMap.StopCleaner()    
```


#### Restarting the cleaner
```go
    timedMap.StartCleaner()
```


#### Restarting the cleaner with a new interval
```go
    interval := time.time.Millisecond * 500

    timedMap.RestartCleanerWithInterval(interval)
```


#### CLEAN.. NOW !
```go
    timedMap.CleanNow()
```
This will start cleaning expired values regardless of the cleaning interval,
It will run on the main goroutine so it a blocking operation. 
