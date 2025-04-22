package tests

/*
func TestSendMail1(t *testing.T) {

	defer monkey.UnpatchAll()
	monkey.Patch((*mail.Email).Send, func(email *mail.Email, client *mail.SMTPClient) error {
		return nil
	})

	monkey.Patch((*mail.Email).Send, func(client *mail.SMTPClient) error {
		return nil
	})

	mailSubject := "test mail subject"
	fromName := "test sender"
	fromEmail := "test@email"
	toName := "test reciever"
	toEmail := "test reciever"
	mailBody := "test reciever"

	actual := handlers.SendEmail(mailSubject, fromName, fromEmail, toName, toEmail, mailBody, "", "")

	if actual != nil {
		t.Errorf("actual %q, expected no error or nil err val", actual)
	}
}
*/

/*
func TestSendMail2(t *testing.T) {

	mailSubject := "test mail subject"
	fromName := "test sender"
	fromEmail := "test@email"
	toName := "test reciever"
	toEmail := "test reciever"
	mailBody := "test reciever"

	// Establish connections
	broker := connection.NatsConnection()
	logrus.Info("Nats connected: ", broker.IsConnected())

	// get NATS subject for email sending service
	natsSubject := os.Getenv("NATS_SUBJECT")

	defer broker.Close()

	ec, err := nats.NewEncodedConn(broker, nats.JSON_ENCODER)
	if err != nil {
		logrus.Fatal(err)
	}
	defer ec.Close()

	// Publish the message
	if err := ec.Publish(natsSubject,
		&model.EmailData{MailSubject: mailSubject, FromName: fromName, FromEmail: fromEmail, ToName: toName, ToEmail: toEmail, MailBody: mailBody}); err != nil {
		logrus.Fatal(err)
	}

	if err != nil {
		t.Error("Error: ", err)
	}
}
*/
