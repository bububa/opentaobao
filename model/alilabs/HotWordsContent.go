package alilabs

// HotWordsContent 结构体
type HotWordsContent struct {
	// 热词列表
	Words []string `json:"words,omitempty" xml:"words>string,omitempty"`
}
