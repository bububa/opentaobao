package qianniu

// PicAuditParam 结构体
type PicAuditParam struct {
	// 图片链接列表
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
