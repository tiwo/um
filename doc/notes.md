* Goal of this project: Learn Go
* I will implement a virtual Universal Machine without looking into previous implementations
* I will take notes in this document, in chronological order, adding at the bottom
* I do not at this time feel like solving the programming problems

### Physical Specs:
* unsigned 32bit words "platters with room for 32 small marks"
* eight registers
* a number of arrays of platters, referenced by a 1-platter pointer
* one of those arrays, the 0-array is referenced by 0 and stores "the program"
* 8-bit standard input and output

### Reading files into "sand" arrays

* binary.[Read](https://pkg.go.dev/encoding/binary#Read)(r io.Reader, order ByteOrder, data interface{}) error

But "data must be a pointer to a fixed-size value or a slice of fixed-size values," maybe I should accept standard input?

* (b *bufio.Reader) [`Read`](https://pkg.go.dev/bufio#Reader.Read)(p []byte) (n int, err error)

reads bytes into p

* io.[ReadFull](https://pkg.go.dev/io#ReadFull)(r Reader, buf []byte) (n int, err error)

reads into a buffer as well.

I'll try the last one.