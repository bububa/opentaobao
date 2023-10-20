package promotion

// AlibabaasrdataservicepromotionrulequeryData 结构体
type AlibabaasrdataservicepromotionrulequeryData struct {
	// 兑换详情列表
	DetailList []Detaillist `json:"detail_list,omitempty" xml:"detail_list>detaillist,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// poskey
	PosKey int64 `json:"pos_key,omitempty" xml:"pos_key,omitempty"`
}
