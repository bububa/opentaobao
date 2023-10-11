package qianniu

// PicAuditResult 结构体
type PicAuditResult struct {
	// 经审核的图片链接列表
	AuditedUrls []string `json:"audited_urls,omitempty" xml:"audited_urls>string,omitempty"`
	// 不合法的图片链接序号
	IllegalIndex []string `json:"illegal_index,omitempty" xml:"illegal_index>string,omitempty"`
}
