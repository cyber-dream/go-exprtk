package exprtk

// #cgo CXXFLAGS: -flto -fuse-linker-plugin -std=c++11
// #cgo LDFLAGS: -L.
// #include "exprtkwrapper.h"
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

// GoExprtk ...Exprtk Structure
type GoExprtk struct {
	exprtk C.exprtkWrapper
}

// NewExprtk ... Creates a new object
func NewExprtk() GoExprtk {
	var obj GoExprtk
	obj.exprtk = C.exprtkWrapperInit()
	return obj
}

// SetExpression ... Sets an Expression
func (obj GoExprtk) SetExpression(expr string) {
	str := C.CString(expr)

	C.setExpressionString(obj.exprtk, str)

	C.free(unsafe.Pointer(str))
}

// AddDoubleVariable ... Adds variable to the expression
func (obj GoExprtk) AddDoubleVariable(x string) {
	str := C.CString(x)

	C.addDoubleVariable(obj.exprtk, str)

	C.free(unsafe.Pointer(str))
}

// AddStringVariable ... Adds variable to the expression
func (obj GoExprtk) AddStringVariable(x string) {
	str := C.CString(x)

	C.addStringVariable(obj.exprtk, str)

	C.free(unsafe.Pointer(str))
}

// AddVectorVariable ... Adds variable to the expression
func (obj GoExprtk) AddVectorVariable(x string) {
	str := C.CString(x)

	C.addVectorVariable(obj.exprtk, str)

	defer C.free(unsafe.Pointer(str))
}

// SetDoubleVariableValue ... Sets value to the variable
func (obj GoExprtk) SetDoubleVariableValue(varName string, val float64) {
	cName := C.CString(varName)

	C.setDoubleVariableValue(obj.exprtk, cName, C.double(val))

	C.free(unsafe.Pointer(cName))
}

// SetStringVariableValue ... Sets value to the variable
func (obj GoExprtk) SetStringVariableValue(varName string, val string) {
	cName := C.CString(varName)
	cVal := C.CString(val)

	C.setStringVariableValue(obj.exprtk, cName, cVal)

	C.free(unsafe.Pointer(cName))
	C.free(unsafe.Pointer(cVal))
}

// SetVectorVariableValue ... Sets value to the variable
func (obj GoExprtk) SetVectorVariableValue(varName string, val []float64) {
	arr := make([]C.double, 0)
	for i := 0; i < len(val); i++ {
		arr = append(arr, C.double(val[i]))
	}
	firstValue := &(arr[0])
	var arrayLength C.int = C.int(len(arr))

	cName := C.CString(varName)

	C.setVectorVariableValue(obj.exprtk, cName, firstValue, arrayLength)

	C.free(unsafe.Pointer(cName))
}

// CompileExpression ... Compiles the Expression
func (obj GoExprtk) CompileExpression() error {
	value := C.compileExpression(obj.exprtk)
	if value == 0 {
		return errors.New("failed to compile the expression")
	}
	return nil
}

// GetEvaluatedValue ... Returns the evaluated value
func (obj GoExprtk) GetEvaluatedValue() float64 {
	return float64(C.getEvaluatedValue(obj.exprtk))
}

// Delete ... Destroys the created object and releases the memory
func (obj GoExprtk) Delete() {
	C.deleteExprtk(obj.exprtk)
}

// CollectVariables Returns names of all variables in Expression
func (obj GoExprtk) CollectVariables() []string {
	arr := C.collectVariables(obj.exprtk)
	defer C.freeVariableNamesArray(arr)

	arrSlice := make([]string, 0, arr.size)

	varDataSlice := unsafe.Slice(arr.data, arr.size)

	for idx := range varDataSlice {
		arrSlice = append(arrSlice, C.GoString(varDataSlice[idx]))
	}

	return arrSlice
}
