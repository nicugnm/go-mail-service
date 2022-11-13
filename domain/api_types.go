package domain

type EmailTemplate struct {
	FromMail string `json:"from_mail"`
	ToMail   string `json:"to_mail"`
	CC       string `json:"mail_cc"`
	body     string `json:"body"`
}
