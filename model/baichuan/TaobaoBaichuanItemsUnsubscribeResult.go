package baichuan

// TaobaobaichuanitemsunsubscribeResult 结构体
type TaobaobaichuanitemsunsubscribeResult struct {
	// 返回按resultCode分为多个返回部分
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}
