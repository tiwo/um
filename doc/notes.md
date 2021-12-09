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