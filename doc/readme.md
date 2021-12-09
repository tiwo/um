### Original [Instructions](http://www.boundvariable.org/rulesfaq.shtml):

- [ ] Register to receive a decryption key.
- [ ] Download the codex and the UM specification.
- [ ] Implement the UM and run the codex.
- [ ] Supply the decryption key you received from us.
- [ ] Obtain publications and submit them.


### Contest materials:

(from the [original task description](http://www.boundvariable.org/task.shtml))


#### UM Specification `um-spec.txt`

Specification for the Universal Machine. Text format.


#### Codex `codex.umz (VOLUME ID 9)`

MD5 hash	e328209bd65ade420371d7bd87b88e4f
SHA-1 hash	088ac79d311db02d9823def598e48f2f8723e98a
Decryption key	(\b.bb)(\v.vv)06FHPVboundvarHRAk


#### SANDmark `sandmark.umz`

Benchmark for the Universal Machine. Expected output.

MD5 hash	1c604d454de05d04afdabd2c63fb27fb
SHA-1 hash	c2ee087aa661e81407fbcf0d9d7e503aff9b268e


#### Reference Implementation `um.um` (1024 bytes)

CMU Reference implementation of the Universal Machine.

This implementation supports all UM programs, including uncompressed .um files and self-decompressing .umz files. To use this implementation to run a UM binary called c.um, simply concatenate the two files together:

    cat um.um c.um > cmu.um

The resulting binary can be run in any compliant universal machine implementation, including itself.