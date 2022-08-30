package einvoice

// TaxOptimizationBillDownloadUrlResultDto 结构体
type TaxOptimizationBillDownloadUrlResultDto struct {
	// 账单文件的下载地址，请求成功后20s内有效
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
}
