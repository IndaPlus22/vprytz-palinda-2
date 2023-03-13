# Task 1 - Debugging Concurrent Programs

In this task, you will be provided with two buggy programs. For each program,
you should:

1. Explain what is wrong with the code.
2. Fix the bug.
3. Explain why your solution fixes the bug.

## Buggy program 1

[src/bug01.go](src/bug01.go)

**Answer**: The program is buggy because the `main` function declares an unbuffered channel `ch`, meaning that the sender is stuck waiting for a reciever to read from the channel. My solution to the bug is to make the channel buffered, so that the sender is not blocked.

## Buggy program 2

[src/bug02.go](src/bug02.go)

**Answer**: The program is buggy because the `main` function (thread) exits before the `Print` function has finished reading from the channel. My solution to the bug is to implement a `WaitGroup` to ensure that the `Print` function has finished reading from the channel before the main thread exits.
