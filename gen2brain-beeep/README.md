# Library gen2brain/beeep
[beeep](https://github.com/gen2brain/beeep) provides a cross-platform library for sending desktop notifications, alerts and beeps.

## Examples
![](print/welcome.png "Notification of welcome")

```go
var err error
if err = beeep.Notify("Hello, World!", "Message body", "assets/hands.png"); err != nil {
  panic(err)
}
```

![](print/schedule.png "Notification of Schedule accepted")

```go
var err error
if err = beeep.Alert("Schedule accepted", "Schedule accepted successfully", "assets/schedule.png"); err != nil {
  panic(err)
}
```

