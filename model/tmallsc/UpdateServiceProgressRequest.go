package tmallsc

// UpdateServiceProgressRequest 结构体
type UpdateServiceProgressRequest struct {
	// 图片地址回传集合
	PicUrlList []string `json:"pic_url_list,omitempty" xml:"pic_url_list>string,omitempty"`
	// 服务描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 服务动作
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 扩展信息回传备注
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}
