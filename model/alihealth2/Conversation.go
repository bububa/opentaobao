package alihealth2

// Conversation 结构体
type Conversation struct {
	// PATIENT(&#34;患者&#34;),     DOCTOR(&#34;医生&#34;),     SYSTEM(&#34;系统&#34;)
	Role string `json:"role,omitempty" xml:"role,omitempty"`
	// YYYY-MM-DD HH:mm:ss格式的时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// TEXT(&#34;文本&#34;),     IMG(&#34;图片&#34;),     VOICE(&#34;语音&#34;)
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 聊天内容，如果是图片或者语音，需要通过base64编码为String后传入。
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
