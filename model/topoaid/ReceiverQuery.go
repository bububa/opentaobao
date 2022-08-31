package topoaid

// ReceiverQuery 结构体
type ReceiverQuery struct {
	// 交易订单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 解密场景编号。不同场景，解密策略不同。请根据产品功能选择相应的场景编号。可选的场景：1001(顺丰电子面单发货)、1002(4通一达电子面单发货)、1003(EMS电子面单发货)、1004(其他电子面单发货)、1005(线下门店发货)、1006(手工单发货)、1007(代发货)、2001(客户售后服务)、2002(客户关怀)，&lt;a href=&#34;https://open.taobao.com/doc.htm?docId=120186&amp;docType=1&#34; target=&#34;_blank&#34;&gt;详情点击&lt;/a&gt;
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 隐私号绑定天数
	SecretNoDays int64 `json:"secret_no_days,omitempty" xml:"secret_no_days,omitempty"`
}
