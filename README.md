# krtime 

[![GoDoc](https://godoc.org/github.com/getogrand/krtime?status.svg)](https://godoc.org/github.com/getogrand/krtime)

Package krtime provides helper functions that ease to use time.Time in Korea.

No more `seoul, err := time.LoadLocation("Asia/Seoul"); if err != nil { ... }; t := time.Now().In(seoul)`.

## Why should use this package?

So much boilerplate code is required when you want to display or manipulate time in KST.

```go
seoul, err := time.LoadLocation("Asia/Seoul")
if err != nil {
  panic("load location %q: %v", "Asia/Seoul", err)
}

now := time.Now().In(seoul)
```

No fun. Too verbose. Can't remember timezone string.

## Installation

Run `go get github.com/getogrand/krtime`.

## Usage

### Get Current Time as KST

```go
import "github.com/getogrand/krtime"

now := krtime.Now()
```

### Convert A Time to KST Time

```go
import "time"
import "github.com/getogrand/krtime"

someTime := time.Now().After(1 * time.Hour)
nowkr := krtime.New(someTime) // It returns pure time.Time
```

### Work with jinzhu's [now](https://github.com/jinzhu/now) package

```go
import "github.com/jinzhu/now"
import "github.com/getogrand/krtime"

todayEnd := now.New(krtime.Now()).EndOfDay()
```

### Work with any other library compatible with time.Time

```go
import "github.com/jinzhu/now"
import "github.com/cactus/gostrftime"
import "github.com/hey-beauty/server-api-up/utils/krtime"

todayEnd := now.New(krtime.Now()).EndOfDay()
timestr := gostrftime.Format("%m월 %d일 %l시%M분 %p", krtime.New(todayEnd))
```
