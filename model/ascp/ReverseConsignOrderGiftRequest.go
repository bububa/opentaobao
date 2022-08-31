package ascp

// ReverseConsignOrderGiftRequest 结构体
type ReverseConsignOrderGiftRequest struct {
	// 赠品列表
	GiftOrders []GiftOrder `json:"gift_orders,omitempty" xml:"gift_orders>gift_order,omitempty"`
	// 分销类型（不传代表非分销场景） 枚举值： (1, &#34;直营和分销全参与&#34;),  (2, &#34;仅直营参与&#34;),  (3, &#34;仅分销参与&#34;),  (4, &#34;指定分销店铺&#34;);
	DistributionType string `json:"distribution_type,omitempty" xml:"distribution_type,omitempty"`
	// 订单来源平台编码,TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、MIA=蜜芽、GW=商家官网、CT=村淘、YJWD=云集微店、PDD=拼多多、 DYXD=抖音小店、KS=快手、OTHERS=其他
	SourcePlatformCode string `json:"source_platform_code,omitempty" xml:"source_platform_code,omitempty"`
	// 交易主单
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 翱象发货单号
	ConsignOrderCode string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单产品类型,必填，不传默认0，枚举值：  (0, &#34;全量产品类型&#34;), (1,&#34;普通产品类型(代表除预售下沉提前发货产品类型以外的类型）&#34;), (2, &#34;预售下沉提前发货产品类型&#34;), (3, &#34;普通预售产品类型&#34;);
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 卖家留言
	SellerMessage string `json:"seller_message,omitempty" xml:"seller_message,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 交易下单时间(毫秒)
	TradeCreateTime int64 `json:"trade_create_time,omitempty" xml:"trade_create_time,omitempty"`
	// 交易支付时间(毫秒)
	TradePayTime int64 `json:"trade_pay_time,omitempty" xml:"trade_pay_time,omitempty"`
	// 订单金额：支付金额-退款金额,单位:分，如:20007，表示:20元7分
	OrderAmount int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// 支付金额，单位:分，如:20007，表示:20元7分
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 收件人结构体
	ReceiverInfo *ReceiverInfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
