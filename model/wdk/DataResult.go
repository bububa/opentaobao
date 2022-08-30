package wdk

// DataResult 结构体
type DataResult struct {
	// 取消详情列表
	Datas []SubOrderReturn `json:"datas,omitempty" xml:"datas>sub_order_return,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
