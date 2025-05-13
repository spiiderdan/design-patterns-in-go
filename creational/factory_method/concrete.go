package factory_method

import "fmt"

func (e *EmailNotifier) Send() {
	fmt.Printf("[Email] To: %s — %s\n", e.Notifier.recipient, e.Notifier.message)
}


func (e *SmsNotifier) Send() {
	fmt.Printf("[SMS] To: %s — %s\n", e.recipient, e.message)
}