package trycatch

import (
	"errors"
	"reflect"
)

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
	errType reflect.Type
	handler reflect.Value
}

// TryCatch the struct.
type TryCatch struct {
	f       func()
	catches []catch
	finally func()
}

// Try gets the function with expected error.
func Try(f func()) *TryCatch {
	return &TryCatch{f, []catch{}, nil}
}

// Catch method
func (tryCatch *TryCatch) Catch(fn interface{}) *TryCatch {
	fnV := reflect.ValueOf(fn)
	if fnV.Kind() != reflect.Func {
		panic(errors.New(".Catch: expected function"))
	}
	fnT := fnV.Type()
	if fnT.NumIn() != 1 {
		panic(errors.New(".Catch: expected function to accept 1 argument"))
	}

	tryCatch.catches = append(tryCatch.catches, catch{fnT.In(0), fnV})
	return tryCatch
}

// Finally method
func (tryCatch *TryCatch) Finally(f func()) *TryCatch {
	tryCatch.finally = f
	return tryCatch
}

// Do runs the Try func, apply catch and finally
func (tryCatch *TryCatch) Do() {

	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(IError)
			if ok {
				var tryCatchErrorType = reflect.TypeOf(r)
				var catched = false

				for _, catcher := range tryCatch.catches {
					if tryCatchErrorType.AssignableTo(catcher.errType) {
						catcher.handler.Call([]reflect.Value{reflect.ValueOf(r)})
						catched = true
						break
					}

					if catcher.errType.Kind() == reflect.Interface && tryCatchErrorType.Implements(catcher.errType) {
						catcher.handler.Call([]reflect.Value{reflect.ValueOf(r)})
						catched = true
						break
					}
				}

				if !catched {
					panic(r)
				}
			}
		}

		if tryCatch.finally != nil {
			tryCatch.finally()
		}
	}()

	tryCatch.f()
}

// RaiseError method
func RaiseError(err *Error) {
	panic(err)
}
