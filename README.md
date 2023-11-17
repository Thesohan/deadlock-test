### Program execution:
#### Command:
 `deadlock git:(main) go run deadlock.go`

#### Output:                                            
```Hello World
hi
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.tickerCleanUp(...)
        /Users/sohankathait/Desktop/deadlock/deadlock.go:11
main.main()
        /Users/sohankathait/Desktop/deadlock/deadlock.go:36 +0x108

goroutine 34 [chan send]:
main.tickerCleanUp(...)
        /Users/sohankathait/Desktop/deadlock/deadlock.go:11
main.main.func1()
        /Users/sohankathait/Desktop/deadlock/deadlock.go:28 +0x9f
created by main.main
        /Users/sohankathait/Desktop/deadlock/deadlock.go:21 +0xd1
exit status 2
```
**Note:**
1. We will only get the deadlock error if all goroutines are asleep. In case any of the goroutine is running, program will wait indefinitely for the response from the closed channel.
