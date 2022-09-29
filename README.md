# Repo for Canonical challenge

Contains answer to:

* Build a BASH script to run a AMD64 Linux filesystem image using QEMU;
* Build a shread function in golang.

## Shred function

The `shred` function can be used to safely delete files from the file system. Before deleting,
writing the file 3 times with random data should allow more security against attacks that can
read the residual charge on the disk that could reveal old data.
This function can be useful when making sure that temporary sensitive data are securely delete
after use. On the other side, this secure deletion will cause a performance hit due to the multiple
writes necessary. It also force the file system to sync (making sure the data are written on the
medium).
This version also do create a single byte buffer the same size of the file it will delete. It could
be a good solution for small file but for bigger it could create problems. Another approach in that
case would be to have a fixed size buffer and write the file in sequence.

### Instruction

To run the tests:

```
go test
```

To run the shred tool:

```
go run .
```

## Known problems

* In the QEMU problem, the VM should print 'Hello World!' but it doesn't. I have tried to use
`cloud-init` to perform the task but I don't have enough experience with it at the moment to
complete the task in the 2 hours window requested;

