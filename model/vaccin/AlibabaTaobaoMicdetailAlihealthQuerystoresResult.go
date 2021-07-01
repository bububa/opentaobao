package vaccin

// AlibabaTaobaoMicdetailAlihealthQuerystoresResult 结构体
type AlibabaTaobaoMicdetailAlihealthQuerystoresResult struct {
	// 返回说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回代码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	VaccinStoreList []VaccinStoreInfo `json:"vaccin_store_list,omitempty" xml:"vaccin_store_list>vaccin_store_info,omitempty"`
}
