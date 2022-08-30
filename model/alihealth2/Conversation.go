package alihealth2

// Conversation 结构体
type Conversation struct {
	// PATIENT("患者"),     DOCTOR("医生"),     SYSTEM("系统")
	Role string `json:"role,omitempty" xml:"role,omitempty"`
	// YYYY-MM-DD HH:mm:ss格式的时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// TEXT("文本"),     IMG("图片"),     VOICE("语音")
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 聊天内容，如果是图片或者语音，需要通过base64编码为String后传入。
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
