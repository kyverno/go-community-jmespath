package error

import (
	"errors"
	"fmt"
)

// NotFoundError is returned when it is impossible to resolve the AstField.
type NotFoundError struct {
	key string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("Unknown key \"%s\" in path", n.key)
}

func NotFound(key string) NotFoundError {
	return NotFoundError{key}
}

func NotAnInteger(name string, arg string) error {
	return errors.New(formatNotAnInteger(name, arg))
}

func NotAPositiveInteger(name string, arg string) error {
	return errors.New(formatNotAPositiveInteger(name, arg))
}

func NotEnoughArgumentsSupplied(name string, count int, minExpected int, variadic bool) error {
	return errors.New(formatNotEnoughArguments(name, count, minExpected, variadic))
}

func TooManyArgumentsSupplied(name string, count int, maxExpected int) error {
	return errors.New(formatTooManyArguments(name, count, maxExpected))
}

func formatNotAnInteger(name string, arg string) string {
	return fmt.Sprintf("invalid value, the function '%s' expects its '%s' argument to be an integer.", name, arg)
}

func formatNotAPositiveInteger(name string, arg string) string {
	return fmt.Sprintf("invalid value, the function '%s' expects its '%s' argument to be a an integer value greater than or equal to zero.", name, arg)
}

func formatNotEnoughArguments(name string, count int, minExpected int, variadic bool) string {
	more := ""
	only := ""

	if variadic {
		more = "or more "
		only = "only "
	}

	report := fmt.Sprintf("%s%d ", only, count)
	if count == 0 {
		report = "none "
	}

	plural := ""
	if minExpected > 1 {
		plural = "s"
	}

	return fmt.Sprintf(
		"invalid arity, the function '%s' expects %d argument%s %sbut %swere supplied",
		name,
		minExpected,
		plural,
		more,
		report)
}

func formatTooManyArguments(name string, count int, maxExpected int) string {
	return fmt.Sprintf(
		"invalid arity, the function '%s' expects at most %d arguments but %d were supplied",
		name,
		maxExpected,
		count)
}
