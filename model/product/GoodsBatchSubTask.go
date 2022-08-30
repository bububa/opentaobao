package product

// GoodsBatchSubTask 结构体
type GoodsBatchSubTask struct {
	// 子任务状态产生原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 外部商品ID，用于标识外部系统每次提交过来的商品
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 子任务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 子任务id
	SubBatchId int64 `json:"sub_batch_id,omitempty" xml:"sub_batch_id,omitempty"`
	// 交易猫商品ID，如果商品发布失败则为空
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
