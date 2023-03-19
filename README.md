# pasataleo/go-errors

This is my golang errors library.

It contains some common implementations of error functionality wrapped into a 
single library. You can create simple new errors, create slices of errors, and 
wrap errors with additional information.

This library can also embed error codes and generic data within the errors it
creates and returns. This allows users to return additional data that can be 
processed by callers.

## Error Codes

A new type of string is defined by this library: `errors.ErrorCode`.

You can use the generic `errors.ErrorCodeUnknown` if you don't want to use this
functionality. If you do, then you can define custom error codes that users can
process to handle specific error cases. It is recommended to prepend the package
name to any error codes to ensure they are unique. For example, if you wished to
return a specific error code indicating an item was found:

```golang
package my_library

import "github.com/pasataleo/go-errors"

const ErrorCodeNotFound errors.ErrorCode = "MyLibraryErrorCodeNotFound"
```

You could then create an error: `errors.New(nil, ErrorCodeNotFound, "item not found")`.

And users of your library can process a not found error: `errors.Is(err, my_library.ErrorCodeNotFound)`.

## New Errors

You can create new errors using the `errors.New` function and the `errors.Newf`
function. For new errors you can specify the wrapped error as nil, and you must
provide an error code.

## Wrapped Errors

The `errors.New` and `errors.Newf` functions can accept a wrapped error and will
override the error code of the wrapped error, or can set an error code if the
wrapped error was not created by this library.

The `errors.Wrap` and `errors.Wrapf` functions must accept a wrapped error and
do not accept a new error code. Errors created with these functions will use
the error code of the wrapped error.

You can unwrap an error using the `errors.Unwrap` function. This will return any
wrapped errors or nil if nothing was wrapped.

## Multi Errors

You can append errors into a slice using the `errors.Append` function:

```golang
err = errors.Append(err, errors.New(nil, my_library.ErrorCodeNotFound, "item not found"))
```

You can pass any error into the append function as the first argument. If it is
already a multi error then the new errors will simply be added into the existing
multi. If it is `nil` or another kind of error, a new multi error will be 
created and the supplied errors combined.