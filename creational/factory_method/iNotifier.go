package factory_method

type INotifier interface{
	setRecipient (recipient string)
	setMessage (message string)

	getRecipient() string
	getMessage() string

	send()
}
