package openmall

// TopTradeCreateDo 结构体
type TopTradeCreateDo struct {
	// 收货地址的收件人地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 买家来源渠道，可选值 WeiXin（微信渠道），DouYin（抖音）， KuaiShou （快手），Other（其他）
	BuyerChannel string `json:"buyer_channel,omitempty" xml:"buyer_channel,omitempty"`
	// 买家来源渠道对应ID，例如微信的openId
	BuyerChannelId string `json:"buyer_channel_id,omitempty" xml:"buyer_channel_id,omitempty"`
	// 买家留言
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 买家手机号
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	Distributor string `json:"distributor,omitempty" xml:"distributor,omitempty"`
	// 收货地址的区域码，从taobao.areas.get接口获取区域码，文档地址：https://open.taobao.com/api.htm?spm=a219a.7386653.0.0.77e7669aoxOfiV&source=search&docId=59&docType=2
	Divisioncode string `json:"divisioncode,omitempty" xml:"divisioncode,omitempty"`
	// 计划采购金额（单位元），如最终订单成交价与期望金额不同则直接关单采购失败，计算方式：运费+商品成本价
	ExpectOrderFee string `json:"expect_order_fee,omitempty" xml:"expect_order_fee,omitempty"`
	// 收货地址的手机号码
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 收货地址的收件人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 订单的外部订单号，用来防止重复提交。需要以TOP：  appkey_（如：100000_） 开头，最长32位
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 外部订单交易金额（单位元），即消费者在渠道的实付金额
	OuterOrderFee string `json:"outer_order_fee,omitempty" xml:"outer_order_fee,omitempty"`
	// 收货地址的电话号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货地址的邮编，为空或者6位数字的邮编
	Postcode string `json:"postcode,omitempty" xml:"postcode,omitempty"`
	// 创建交易时的物流方式。 具体的值从 taobao.openmall.trade.render 接口获取，邮费0说明为包邮
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// 商品的id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品数量
	Nums int64 `json:"nums,omitempty" xml:"nums,omitempty"`
	// 商品对应的SKUID，无SKU商品传0
	SkuIids int64 `json:"sku_iids,omitempty" xml:"sku_iids,omitempty"`
	// 创建订单传入true的时候，订单支付后会进入半小时的hold单（订单状态为：PAID_FORBID_CONSIGN），此时订单不会发货，用户可以调用Openmall地址修改接口修改订单收货地址，半小时结束后订单自动结束hold单进入发货流程订单状态为（WAIT_SELLER_SEND_GOODS）
	NeedErpHold bool `json:"need_erp_hold,omitempty" xml:"need_erp_hold,omitempty"`
}
