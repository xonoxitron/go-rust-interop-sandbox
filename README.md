## Go and Rust Interop Sandbox

This example demonstrates how to achieve interoperability between Go and Rust by calling a Rust function from Go code. We'll use Go's cgo and Rust's Foreign Function Interface (FFI) to accomplish this.

### Rust Library

First, let's create a Rust library that exports a function.

Create a new file named `hello.rs` with the following content:

```rust
// hello.rs
#[no_mangle]
pub extern "C" fn greet() {
    println!("Hello from Rust!");
}
```

The `greet()` function simply prints a greeting message.

### Building the Rust Library

Compile the Rust code into a shared library using the following command:

```bash
rustc --crate-type cdylib hello.rs
```

This command generates a shared library file named `libhello.{so|dll|dylib}` depending on your platform.

### Go Code

Next, let's create the Go code that will call the Rust function.

Create a new file named `main.go` with the following content:

```go
// main.go
package main

// #cgo LDFLAGS: -L./ -lhello
// void greet();
import "C"

func main() {
    C.greet()
}
```

The Go code imports the `C` package, which is used for interacting with C code through cgo. It also includes a directive `#cgo LDFLAGS` to link against the Rust shared library.

### Building and Running the Code

To build and run the Go code, execute the following commands:

```bash
go build
./main
```

If everything is set up correctly, you should see the following output:

```
Hello from Rust!
```

Congratulations! You have successfully achieved interop between Go and Rust.

### Additional Notes

- Make sure you have both Go and Rust installed on your system before running this example.
- The Rust code exports a function named `greet()`, which is called from the Go code.
- Ensure that the Rust shared library (`libhello.{so|dll|dylib}`) is in the same directory as the Go executable or in a directory specified in the system's library search path.
- The `#[no_mangle]` attribute ensures that the Rust function name is not mangled, making it compatible with C-style function names.
- You can extend this example by passing arguments between Go and Rust or returning values from Rust functions to Go.
- Remember to handle potential errors and edge cases based on your specific requirements.

That's it! You now have a basic understanding of how to achieve Go and Rust interop. Feel free to explore more advanced scenarios and features provided by both languages for seamless integration.
