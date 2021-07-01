package viapi

// AliyunViapiOcrCharacterData 结构体
type AliyunViapiOcrCharacterData struct {
	// 返回识别信息
	Results []AliyunViapiOcrCharacterResult `json:"results,omitempty" xml:"results>aliyun_viapi_ocr_character_result,omitempty"`
}
