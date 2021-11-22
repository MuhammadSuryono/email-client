package models

type ParamSendMessage struct {
	Recipients  	[]string `form:"msisdn" json:"recipients"`
	RecipientsCC 	[]string `form:"message" json:"recipients_cc"`
	Subject 		string `form:"subject" json:"subject"`
	TypeBody 		string `form:"type_body" json:"type_body"`
	Attachment 		AttachmentFile `form:"attachment" json:"attachment"`
	Body 			string `form:"body" json:"body"`
}

type AttachmentFile struct {
	Filename string `form:"filename" json:"filename"`
	Url string `form:"url" json:"url"`
}