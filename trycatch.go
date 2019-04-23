package main

import "fmt"
import "reflect"

var tryCatchError IError

// IError interface
type IError interface {
	GetMessage() string
}

// Error is default IError implementation
type Error struct {
	str string
}

// GetMessage method
func (err Error) GetMessage() string {
	return err.str
}

type catch struct {
	errType string
	f func(err IError)
}

// TryCatch the struct.
type TryCatch struct {
	f func()
	catches []catch
	finally func()
}

// Try method
func Try(f func()) *TryCatch {
    return &TryCatch{f, []catch{}, nil}
}

// Catch method
func (tryCatch *TryCatch) Catch(err string, f func(err IError)) *TryCatch {
	tryCatch.catches = append(tryCatch.catches, catch{err, f})
	return tryCatch
}

// Finally method
func (tryCatch *TryCatch) Finally(f func()) *TryCatch {
	tryCatch.finally = f
	return tryCatch
}

// Do method
func (tryCatch *TryCatch) Do() {

    defer func() {
        if r := recover(); r != nil {
			tryCatchErr, ok := r.(IError)
			if (ok){
				var tryCatchErrorType = reflect.TypeOf(tryCatchErr)

				for _, c := range tryCatch.catches {
					if c.errType == tryCatchErrorType.Name() {
						tryCatchError = nil
						c.f(tryCatchErr)
					}
				}

				if (tryCatch.finally != nil){
					tryCatch.finally()
				}
			}
        }
	}()
	
	tryCatch.f()
}

// RaiseError method
func RaiseError(err *Error) {
	//tryCatchError = err
	panic(err)
}



func main()  {
	Try(func(){
		fmt.Println("Hello world!")
		RaiseError(&Error{"Oh no!"})
	}).
	Catch("Error", func(err IError){
		fmt.Println("Catch found!")
	}).
	Finally(func(){}).
	Do()
}