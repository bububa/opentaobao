package product

// GoodsBatchTaskResultDto 结构体
type GoodsBatchTaskResultDto struct {
	// 商品批次子任务对象集合
	GoodsBatchSubTaskList []GoodsBatchSubTask `json:"goods_batch_sub_task_list,omitempty" xml:"goods_batch_sub_task_list>goods_batch_sub_task,omitempty"`
	// 任务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
