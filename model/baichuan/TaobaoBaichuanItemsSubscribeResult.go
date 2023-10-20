package baichuan

// TaobaobaichuanitemssubscribeResult 结构体
type TaobaobaichuanitemssubscribeResult struct {
	// 按不同的返回码将结果分部分返回
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}
