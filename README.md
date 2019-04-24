# Go try-catch
[![Build Status](https://travis-ci.org/alexpantyukhin/go-try-catch.svg?branch=master
)](https://travis-ci.org/alexpantyukhin/go-try-catch)
[![GoDoc](https://godoc.org/alexpantyukhin/go-try-catch?status.svg)](https://godoc.org/github.com/alexpantyukhin/go-try-catch)

# Motivation
ALL LANGUAGES SHOULD LOOK LIKE C#))) If you are still there then.. don't worry it's a joke) But many a true word is spoken in jest. This repository was made for fun. I don't want to get some lib which will be used in some real systems making this project. My advise: DON'T USE IT IN REAL PROJECTS :).

# Usages
It's possible to try use try-catch:

```go
	Try(func() {
		fmt.Println("Hello Try! Let's get some fun!")
		RaiseError(&Error{"Oh no! something went wrong! Let's get out of here!!!"})
	}).
	Catch(func(err *Error) {
		fmt.Println("Hehe... So good we have Catch there.")
	}).
	Finally(func(){
		fmt.Println("Roll up, time to go home.")
	}).
	Do()
```


# Installation
Just `go get` this repository with the following way:

```
go get github.com/alexpantyukhin/go-try-catch
```