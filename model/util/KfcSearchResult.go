package util

// KfcSearchResult 结构体
type KfcSearchResult struct {
	// 匹配到的关键词的等级，值为null，或为A、B、C、D。<br/>当匹配不到关键词时，值为null，否则值为A、B、C、D中的一个。<br/>A、B、C、D等级按严重程度从高至低排列。
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 过滤后的文本：<br/>当匹配到B等级的词时，文本中的关键词被替换为*号，content即为关键词替换后的文本；<br/>其他情况，content始终为null
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否匹配到关键词,匹配到则为true.
	Matched bool `json:"matched,omitempty" xml:"matched,omitempty"`
}
