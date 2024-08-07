package jmespath

import (
	"github.com/kyverno/go-community-jmespath/pkg/api"
	jperror "github.com/kyverno/go-community-jmespath/pkg/error"
	"github.com/kyverno/go-community-jmespath/pkg/functions"
	"github.com/kyverno/go-community-jmespath/pkg/interpreter"
	"github.com/kyverno/go-community-jmespath/pkg/parsing"
)

// api types

type JMESPath = api.JMESPath

var (
	Compile     = api.Compile
	MustCompile = api.MustCompile
	Search      = api.Search
)

// interpreter types

type Option = interpreter.Option

var WithFunctionCaller = interpreter.WithFunctionCaller

// parsing types

type SyntaxError = parsing.SyntaxError

var NewParser = parsing.NewParser

// function types

type (
	JpFunction    = functions.JpFunction
	JpType        = functions.JpType
	FunctionEntry = functions.FunctionEntry
	ArgSpec       = functions.ArgSpec
	ExpRef        = functions.ExpRef
)

const (
	JpNumber      = functions.JpNumber
	JpString      = functions.JpString
	JpArray       = functions.JpArray
	JpObject      = functions.JpObject
	JpArrayArray  = functions.JpArrayArray
	JpArrayNumber = functions.JpArrayNumber
	JpArrayString = functions.JpArrayString
	JpExpref      = functions.JpExpref
	JpAny         = functions.JpAny
)

// error types

type NotFoundError = jperror.NotFoundError
