## Overview

Worktime file is designed to convert a time string from 12H format to 24H.
The file utilizes the `time` package in Go to parse and format the time
and `flag` package to parse user input


## Usage

Clone the repository and run the program using a Go compiler:

```
git clone https://github.com/c1kzy/Timers-and-signals.git
cd Timers-and-signals
go run main.go (your time here. Example: "12:04:30AM")
```

## Unit Tests
To run unit tests, follow the steps below:

```
git clone https://github.com/c1kzy/Timers-and-signals.git
cd Timers-and-signals
go run test
```

Additionally, you can use -v flag to get detailed information about each test

```
go run test -v
```
