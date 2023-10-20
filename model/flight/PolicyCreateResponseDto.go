package flight

// PolicyCreateResponseDto 结构体
type PolicyCreateResponseDto struct {
	// 错误信息
	ErrorPolicyList []PolicyResponseDto `json:"error_policy_list,omitempty" xml:"error_policy_list>policy_response_dto,omitempty"`
	// 修改日期
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建日期
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 操作人昵称
	AgentSubNick string `json:"agent_sub_nick,omitempty" xml:"agent_sub_nick,omitempty"`
	// 失败excel的地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 属性信息
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 创建结果
	CreateResult *AlitrippolicynormalcompressionuploadResultDto `json:"create_result,omitempty" xml:"create_result,omitempty"`
	// 删除结果
	DeleteResult *AlitrippolicynormalcompressionuploadResultDto `json:"delete_result,omitempty" xml:"delete_result,omitempty"`
	// 政策进度类型
	PolicyProcessType int64 `json:"policy_process_type,omitempty" xml:"policy_process_type,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 成功数量
	SuccessAmount int64 `json:"success_amount,omitempty" xml:"success_amount,omitempty"`
	// 操作人id
	AgentSubId int64 `json:"agent_sub_id,omitempty" xml:"agent_sub_id,omitempty"`
	// 总数量
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 政策类型
	PolicyType int64 `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 失败数量
	FailAmount int64 `json:"fail_amount,omitempty" xml:"fail_amount,omitempty"`
	// 任务状态：0，未开始；1，执行中；2，上传执行完成；3，删除执行完成；4，任务执行完成
	TaskStatus int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
