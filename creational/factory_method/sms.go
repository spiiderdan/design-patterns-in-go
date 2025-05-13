package factory_method

type SmsNotifier struct{
	Notifier
}

func NewSmsNotifier() INotifier {
	return &struct {
		INotifier
		*SmsNotifier
	}{
		SmsNotifier: &SmsNotifier{
			Notifier: Notifier{
				recipient: "",
				message:   "",
			},
		},
	}
}	