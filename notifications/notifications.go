package notifications

import "fmt"

type Notifier interface {
    Notify(message string)
}

type EmailNotifier struct {
    EmailAddress string
}

func (e EmailNotifier) Notify(msg string) {
    fmt.Printf("Отправляю email на %s: \"%s\"\n", e.EmailAddress, msg)
}

type SMSNotifier struct {
    PhoneNumber string
}

func (s SMSNotifier) Notify(msg string) {
    fmt.Printf("Отправляю SMS на номер %s: \"%s\"\n", s.PhoneNumber, msg)
}