package viapi

// AliyunViapiOcrCharacterResult 
type AliyunViapiOcrCharacterResult struct {
    // 文字区域
    TextRectangle   *TextRectangle `json:"text_rectangle,omitempty" xml:"text_rectangle,omitempty"`
    // 文字区域概率，概率值的范围为[0, 1]
    Probability   int64 `json:"probability,omitempty" xml:"probability,omitempty"`
    // 文字内容
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
}