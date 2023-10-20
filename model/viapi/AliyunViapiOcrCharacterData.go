package viapi

// AliyunviapiocrcharacterData 结构体
type AliyunviapiocrcharacterData struct {
	// 返回识别信息
	Results []AliyunviapiocrcharacterResult `json:"results,omitempty" xml:"results>aliyunviapiocrcharacter_result,omitempty"`
}
