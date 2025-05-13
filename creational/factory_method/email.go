package factory_method

type EmailNotifier struct{
	Notifier
}

func NewEmailNotifier() INotifier {
	return &struct {
		INotifier
		*EmailNotifier
	}{
		EmailNotifier: &EmailNotifier{
			Notifier: Notifier{
				recipient: "",
				message:   "",
			},
		},
	}
}	