package einvoice

// InvoiceApplyDtlQueryDTO 
type InvoiceApplyDtlQueryDTO struct {
    // 中台发票申请ID，由中台生成。字母或数字组成。  可用于查询发票申请的详情。
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 是否需要生成发票板式文件的下载链接。默认为：false不生成。  调用方请根据使用场景而定。  true: 生成新的发票下载链接，拆单情况下生成多张发票链接响应时间较长，建议设置超时时间为6s.   false: 不生成下载链接，查询响应时间更快。调用方如果不需要发票链接，或者有对接查询发票详情接口时，此处建议传false
    NeedDownloadUrl   bool `json:"need_download_url,omitempty" xml:"need_download_url,omitempty"`
}
