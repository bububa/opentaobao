package miniapp

// SellerDeliveryTaskDto 结构体
type SellerDeliveryTaskDto struct {
	// 扩展参数
	ExtProperties string `json:"ext_properties,omitempty" xml:"ext_properties,omitempty"`
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 创建渠道。创建渠道。&#34;PRIVATE_DOMAIN&#34;, &#34;私域工作台&#34;；&#34;ISV_SCRM&#34;, &#34;服务商SCRM&#34;；&#34;OTHER&#34;, &#34;其他&#34;
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 子渠道场景，比如新品发布、自定义种草
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 任务名称
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 服务商的out id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 任务状态。&#34;INIT&#34;, &#34;未执行&#34;; &#34;TEMP&#34;, &#34;草稿态&#34;; &#34;SENDING&#34;, &#34;执行中&#34;; &#34;DONE&#34;, &#34;已执行&#34;; &#34;CANCEL&#34;, &#34;已终止&#34;
	TaskStatus string `json:"task_status,omitempty" xml:"task_status,omitempty"`
	// 策略id
	StrategyId int64 `json:"strategy_id,omitempty" xml:"strategy_id,omitempty"`
}
