package baichuan

// TaobaobaichuanitemsubscribeResult 结构体
type TaobaobaichuanitemsubscribeResult struct {
	// 返回的列表
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}
