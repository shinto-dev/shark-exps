package combinederrors

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
)

func Example() {
	//Refer: https://pkg.go.dev/go.uber.org/multierr

	// Errors can be combined with the use of
	//multierr.Combine(
	//	reader.Close(),
	//	writer.Close(),
	//	conn.Close(),
	//)

	err := multierr.Append(errors.New("some error"), errors.New("another error"))
	fmt.Println(err)

	//The underlying list of errors for a returned error object may be retrieved with the Errors function.
	errs := multierr.Errors(err)
	if len(errs) > 0 {
		fmt.Println("The following errors occurred:", errs)
	}

	// There are more helpful functions, check the documentation
}
