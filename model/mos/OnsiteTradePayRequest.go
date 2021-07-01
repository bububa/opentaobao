package mos

// OnsiteTradePayRequest 结构体
type OnsiteTradePayRequest struct {
	// 商户支付流水号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 支付授权码。 消费者喵街中的“付款码”信息
	AuthCode string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`
	// 订单总金额，单位为分，取值范围[0,10000000000]
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 不可打折金额。不参与优惠计算的金额，单位为分，取值范围[0,10000000000]。取值不能超过订单总金额
	UndiscountableAmount string `json:"undiscountable_amount,omitempty" xml:"undiscountable_amount,omitempty"`
	// 允许使用的资金渠道。如果不指定，则可以使用所有资金渠道；如果存在多个资金渠道，以逗号分隔，取值：mj_vcard 和 alipay。mj_vcard代表喵街储值卡，alipay代表支付宝。此处设定的顺序代表了资金渠道出现的顺序。在实际支付时，还会判断其它条件控制资金渠道（如使用储值权益，即使在创建订单时指定资金渠道为支付宝，也必须使用储值卡支付）
	AllowablePayChannels string `json:"allowable_pay_channels,omitempty" xml:"allowable_pay_channels,omitempty"`
	// 是否自动完成消费者对订单的确认。如果是，则无需消费者确认环节，直接进入待付款。取值：Y、N。缺省值为N
	BuyerAutoConfirm string `json:"buyer_auto_confirm,omitempty" xml:"buyer_auto_confirm,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 订单描述。对交易或商品的描述
	Body string `json:"body,omitempty" xml:"body,omitempty"`
	// 商品明细列表。订单包含的商品列表信息
	GoodsDetailList []GoodsDetail `json:"goods_detail_list,omitempty" xml:"goods_detail_list>goods_detail,omitempty"`
	// 商户门店编号的类型。取值：miaojie和out。如果取值为miaojie，则store_id的取值为商户门店在喵街中的编号；如果取值为out，则store_id的取值为商户自己的编号
	StoreIdType string `json:"store_id_type,omitempty" xml:"store_id_type,omitempty"`
	// 商户门店编号。可以是喵街内的商户门店ID，也可以是商户系统内自己的门店ID，其取值的含义由store_id_type定义
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商户操作员编号
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 商户机具终端编号
	TerminalId string `json:"terminal_id,omitempty" xml:"terminal_id,omitempty"`
	// 支付超时时间。该笔订单允许的最晚付款时间，逾期将关闭交易。格式为：yyyy-MM-dd HH:mm:ss。时区为GMT+8。 线下条码支付，建议传入当前时间往后10分钟的时间点
	TimeExpire string `json:"time_expire,omitempty" xml:"time_expire,omitempty"`
	// 业务扩展参数，json格式。目前需要的扩展参数名有：memberPayerSame：是否需要校验会员与支付者是否同人。
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 商户收银支付宝账号
	StoreAlipayAccount string `json:"store_alipay_account,omitempty" xml:"store_alipay_account,omitempty"`
	// 场景流水号，会员识别时生成。
	SceneNo string `json:"scene_no,omitempty" xml:"scene_no,omitempty"`
	// 会员账号ID
	MemberAccountId string `json:"member_account_id,omitempty" xml:"member_account_id,omitempty"`
	// 会员手机号
	MemberMobile string `json:"member_mobile,omitempty" xml:"member_mobile,omitempty"`
	// 支付场景，条码支付：bar_code，刷脸支付：security_code
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
}
