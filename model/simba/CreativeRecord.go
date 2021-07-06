package simba

// CreativeRecord 结构体
type CreativeRecord struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创意标题，最多20个汉字
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 修改前创意标题
	OldTitle string `json:"old_title,omitempty" xml:"old_title,omitempty"`
	// 创意图片地址，必须是推广组对应商品的图片之一
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 修改前创意图片地址
	OldImgUrl string `json:"old_img_url,omitempty" xml:"old_img_url,omitempty"`
	// 审核状态： audit_wait-待审核；audit_pass-审核通过(上线)；audit_reject-审核拒绝；默认为audit_pass。
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 审核拒绝原因描述
	AuditDesc string `json:"audit_desc,omitempty" xml:"audit_desc,omitempty"`
	// 创意被修改的时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 系统记录最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 关键词id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
}
