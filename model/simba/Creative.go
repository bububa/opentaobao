package simba

// Creative 结构体
type Creative struct {
	// 主人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 创意标题，最多30个汉字
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 创意图片地址，必须是推广组对应商品的图片之一
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 审核状态： audit_wait-待审核；audit_pass-审核通过(上线)；audit_reject-审核拒绝；默认为audit_pass。
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 审核拒绝原因描述
	AuditDesc string `json:"audit_desc,omitempty" xml:"audit_desc,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 最后修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 使用副图作为创意的原始副图图片URL后缀
	SecondImgOriginUrl string `json:"second_img_origin_url,omitempty" xml:"second_img_origin_url,omitempty"`
	// 广审批准文号
	AdExaminationCode string `json:"ad_examination_code,omitempty" xml:"ad_examination_code,omitempty"`
	//
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 推广计划Id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 推广组id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意图片类型,1-主图,2-副图,3-自定义图片
	ImgType int64 `json:"img_type,omitempty" xml:"img_type,omitempty"`
	//
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}
