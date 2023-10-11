package xhotelitem

// HotelXitemDo 结构体
type HotelXitemDo struct {
	// 酒+X 图片格式化信息
	Pictures []HotelXItemPicture `json:"pictures,omitempty" xml:"pictures>hotel_x_item_picture,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 外部code
	OutXCode string `json:"out_x_code,omitempty" xml:"out_x_code,omitempty"`
	// 子类型code
	SubTypeCode string `json:"sub_type_code,omitempty" xml:"sub_type_code,omitempty"`
	// 外部酒店code
	OutHid string `json:"out_hid,omitempty" xml:"out_hid,omitempty"`
	// 元素类型简写
	ShortName string `json:"short_name,omitempty" xml:"short_name,omitempty"`
	// 服务时间段
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 商品使用说明
	ItemDesc string `json:"item_desc,omitempty" xml:"item_desc,omitempty"`
	// 审核拒绝原因
	AuditRejectReason string `json:"audit_reject_reason,omitempty" xml:"audit_reject_reason,omitempty"`
	// 详细信息json字符串
	FeatureDetail string `json:"feature_detail,omitempty" xml:"feature_detail,omitempty"`
	// 商品价值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 状态是否生效0 失效, 1生效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 附加产品使用维度   1:每间房维度 2:每间夜维度
	DimensionType int64 `json:"dimension_type,omitempty" xml:"dimension_type,omitempty"`
	//  审核状态-1：拒绝，0：审核中，1：审核通过
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
}
