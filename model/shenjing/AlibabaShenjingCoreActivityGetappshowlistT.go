package shenjing

// AlibabaShenjingCoreActivityGetappshowlistT 结构体
type AlibabaShenjingCoreActivityGetappshowlistT struct {
	// 获取地址
	FullAdress string `json:"full_adress,omitempty" xml:"full_adress,omitempty"`
	// 活动详情的URL地址
	ActDetailUrl string `json:"act_detail_url,omitempty" xml:"act_detail_url,omitempty"`
	// 活动的内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 发布状态  "draft":"草稿","wait":"待发布",  "underway":"进行中",  "over":"已结束",  "insertingcoil":"已下线"
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 活动展示结束时间
	ViewEndTime string `json:"view_end_time,omitempty" xml:"view_end_time,omitempty"`
	// 活动展示开始时间
	ViewStartTime string `json:"view_start_time,omitempty" xml:"view_start_time,omitempty"`
	// 活动结束时间
	ActEndTime string `json:"act_end_time,omitempty" xml:"act_end_time,omitempty"`
	// 活动开始时间
	ActStartTime string `json:"act_start_time,omitempty" xml:"act_start_time,omitempty"`
	// 物业公司ID
	WyCompanyId int64 `json:"wy_company_id,omitempty" xml:"wy_company_id,omitempty"`
	// 活动主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 宣传图片多个,分隔
	Images string `json:"images,omitempty" xml:"images,omitempty"`
	// 活动创建的时间戳
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 活动名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
