package miniappopen

// DistributionOrderQueryByIdOpenRequest 结构体
type DistributionOrderQueryByIdOpenRequest struct {
	// 投放计划的id列表
	OrderIdList []int64 `json:"order_id_list,omitempty" xml:"order_id_list>int64,omitempty"`
}
