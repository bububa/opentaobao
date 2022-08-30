package topoaid

// SecretNoExtendRequest 结构体
type SecretNoExtendRequest struct {
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 解密场景编号。不同场景，解密策略不同。请根据产品功能选择相应的场景编号。可选的场景：1001(顺丰电子面单发货)、1002(4通一达电子面单发货)、1003(EMS电子面单发货)、1004(其他电子面单发货)、1005(线下门店发货)、1006(手工单发货)、1007(代发货)、2001(客户售后服务)、2002(客户关怀)
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 虚拟号有效期延长天数，发货场景只能填7、15、30
	ExtendDays int64 `json:"extend_days,omitempty" xml:"extend_days,omitempty"`
	// 交易订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
