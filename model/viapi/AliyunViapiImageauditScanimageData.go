package viapi

// AliyunViapiImageauditScanimageData 结构体
type AliyunViapiImageauditScanimageData struct {
	// 图片检测结果
	Results []AliyunViapiImageauditScanimageResult `json:"results,omitempty" xml:"results>aliyun_viapi_imageaudit_scanimage_result,omitempty"`
}
