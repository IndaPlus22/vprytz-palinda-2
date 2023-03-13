# Written answers

## Task 1 - Debugging Concurrent Programs

In this task, you will be provided with two buggy programs. For each program,
you should:

1. Explain what is wrong with the code.
2. Fix the bug.
3. Explain why your solution fixes the bug.

### Buggy program 1

[src/bug01.go](src/bug01.go)

**Answer**: The program is buggy because the `main` function declares an unbuffered channel `ch`, meaning that the sender is stuck waiting for a reciever to read from the channel. My solution to the bug is to make the channel buffered, so that the sender is not blocked.

### Buggy program 2

[src/bug02.go](src/bug02.go)

**Answer**: The program is buggy because the `main` function (thread) exits before the `Print` function has finished reading from the channel. My solution to the bug is to implement a `WaitGroup` to ensure that the `Print` function has finished reading from the channel before the main thread exits.

## Task 2 - Many Senders; Many Receivers

The program [many2many.go](src/many2many.go) contains four
producers that together send 32 strings over a channel. At the
other end there are two consumers that receive the strings.
Describe what happens, and explain why it happens, if you make the
following changes in the program. Try first to reason your way
through, and then test your hypothesis by changing and running the
program.

- What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?

  - **Answer**: The program will crash, since then one of the producers will try to send a message to a closed channel (wich will cause a panic).

- What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?
  - **Answer**: This will also cause a panic, since the channel will be closed before all the producers have finished sending messages. This will cause the producers to try to send messages to a closed channel, which will cause a panic. The fact that one producer is finished is not a guarantee that all producers are finished, which is why this will cause a panic.
- What happens if you remove the statement `close(ch)` completely?
  - **Answer**: In theory, this should work, since the channel will only be closed after `wgp.Wait()` meaning no routines that may try to send messages to the channel will be running. However, not closing the channel might mean that some consumers are not able to finish before the program exits, since the channel will not be closed.
- What happens if you increase the number of consumers from 2 to 4?
  - **Answer**: In theory, program should still work fine but it should complete twice as fast, since there are double the amount of consumers to read from the channel, and the program stimulates that it takes some time to consume data from the channel.
- Can you be sure that all strings are printed before the program
  stops?
  - **Answer**: No, since the program only waits for the producers to finish, and not the consumers. This means that the program might exit before all consumers have finished reading from the channel.

Finally, modify the code by adding a new WaitGroup that waits for
all consumers to finish.
