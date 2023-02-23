package usergrowth2

// PushSendBatchRequest 结构体
type PushSendBatchRequest struct {
	// 批量发送消息元数据请求集合
	BatchSendMetaReqs []PushBatchSendMetaReq `json:"batch_send_meta_reqs,omitempty" xml:"batch_send_meta_reqs>push_batch_send_meta_req,omitempty"`
	// 活动id
	ActId string `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 发送端
	SendClient string `json:"send_client,omitempty" xml:"send_client,omitempty"`
	// 消息模板id
	NotifyId string `json:"notify_id,omitempty" xml:"notify_id,omitempty"`
}
