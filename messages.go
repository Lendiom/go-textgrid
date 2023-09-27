package textgrid

type IncomingMessage struct {
	AccountSID        string `form:"AccountSid"`
	APIVersion        string `form:"ApiVersion"`
	Body              string `form:"Body"`
	From              string `form:"From"`
	MessageSID        string `form:"MessageSid"`
	NumberOfMedia     int    `form:"NumMedia"`
	NumberOfSegments  int    `form:"NumSegments"`
	MediaUrl0         string `form:"MediaUrl0"`
	MediaContentType0 string `form:"MediaContentType0"`
	MediaUrl1         string `form:"MediaUrl1"`
	MediaContentType1 string `form:"MediaContentType1"`
	MediaUrl2         string `form:"MediaUrl2"`
	MediaContentType2 string `form:"MediaContentType2"`
	MediaUrl3         string `form:"MediaUrl3"`
	MediaContentType3 string `form:"MediaContentType3"`
	SmsMessageSid     string `form:"SmsMessageSid"`
	SmsSID            string `form:"SmsSid"`
	SmsStatus         string `form:"SmsStatus"`
	To                string `form:"To"`
}
