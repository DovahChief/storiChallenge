package emailSender

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/model"
)

func SendEmail(ctx context.Context, email string, report model.Report) {
	m := mail.NewV3Mail()
	e := mail.NewEmail("Jose", "pepelimonta@gmail.com")
	m.SetFrom(e)

	m.SetTemplateID("d-24517400f8c04e85a525523715b5fb48")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Example User", email),
	}
	p.AddTos(tos...)
	p.SetDynamicTemplateData("totalBalance", fmt.Sprintf("%.2f", report.TotalBalance))
	p.SetDynamicTemplateData("avgDebitAmount", fmt.Sprintf("%.2f", report.AverageDebitAmount))
	p.SetDynamicTemplateData("avgCreditAmount", fmt.Sprintf("%.2f", report.AverageCreditAmount))
	p.SetDynamicTemplateData("transactionsPerMonth", formatReport(report.TransactionsPerMonth))

	m.AddPersonalizations(p)

	rawDecodedText, err := base64.StdEncoding.DecodeString("U0cuZ3BScjZpS2ZSWkc1dDE0RlpJMk9xUS5ubGxqMzYwLVpSaGVHMkF6Y0pPa2I5R3FGemo0dDNWeXFOcHcxVGRDalFr")
	if err != nil {
		panic(err)
	}

	client := sendgrid.NewSendClient(string(rawDecodedText))
	_, err = client.Send(m)
	if err != nil {
		logger.Errorf(ctx, "Error sending email [%v]", err)
	}
}

func formatReport(tpm map[int]int) string {

	result := ""
	for key, element := range tpm {
		result += time.Month(key).String() + " : " + strconv.Itoa(element) + "\n"
	}
	return result
}
