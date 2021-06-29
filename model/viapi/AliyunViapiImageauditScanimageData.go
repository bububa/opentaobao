package viapi

// AliyunViapiImageauditScanimageData 
type AliyunViapiImageauditScanimageData struct {
    // 图片检测结果
    Results   []AliyunViapiImageauditScanimageResult `json:"results,omitempty" xml:"results>aliyun_viapi_imageaudit_scanimage_result,omitempty"`
}
