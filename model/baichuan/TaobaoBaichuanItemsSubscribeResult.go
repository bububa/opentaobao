package baichuan

// TaobaoBaichuanItemsSubscribeResult 结构体
type TaobaoBaichuanItemsSubscribeResult struct {
	// 按不同的返回码将结果分部分返回
	ResultList []ResultMeta `json:"result_list,omitempty" xml:"result_list>result_meta,omitempty"`
}
