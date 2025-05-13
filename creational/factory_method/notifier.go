package factory_method

type Notifier struct{
	recipient string
	message string
}

func (n *Notifier) setRecipient(recipient string){
	n.recipient = recipient
}

func (n *Notifier) getRecipient() string{
	return n.recipient
}

func (n *Notifier) setMessage(message string){
	n.message = message
}

func (n *Notifier) getMessage() string{
	return n.message
}
