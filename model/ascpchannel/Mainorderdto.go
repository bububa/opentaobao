package ascpchannel

// Mainorderdto 结构体
type Mainorderdto struct {
	// JIT协议ID
	OperationOrderId string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
	// 业务活动代码, 新增：FU010，修改：FU020，停用：FU030
	BizActivityCode string `json:"biz_activity_code,omitempty" xml:"biz_activity_code,omitempty"`
	// 请求唯一号
	OperationCode string `json:"operation_code,omitempty" xml:"operation_code,omitempty"`
	// 供应链原始单据来源平台
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 物流货主ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
