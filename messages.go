package textgrid

type IncomingMessage struct {
	AccountSID       string `form:"AccountSid"`
	APIVersion       string `form:"ApiVersion"`
	Body             string `form:"Body"`
	From             string `form:"From"`
	MessageSID       string `form:"MessageSid"`
	NumberOfMedia    int    `form:"NumMedia"`
	NumberOfSegments int    `form:"NumSegments"`
	SmsMessageSid    string `form:"SmsMessageSid"`
	SmsSID           string `form:"SmsSid"`
	SmsStatus        string `form:"SmsStatus"`
	To               string `form:"To"`
}
