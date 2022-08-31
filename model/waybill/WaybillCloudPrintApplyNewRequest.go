package waybill

// WaybillCloudPrintApplyNewRequest 结构体
type WaybillCloudPrintApplyNewRequest struct {
	// 请求面单信息，数量限制为10
	TradeOrderInfoDtos []TradeOrderInfoDto `json:"trade_order_info_dtos,omitempty" xml:"trade_order_info_dtos>trade_order_info_dto,omitempty"`
	// &lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.8cf9Nj&amp;treeId=17&amp;articleId=105085&amp;docType=1#1&#34;&gt;物流公司Code&lt;/a&gt;，长度小于20
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 目前仅顺丰支持此字段，传入快递产品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 仓code， 仓库WMS系统对接落地配业务，其它场景请不要使用
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 配送资源code， 仓库WMS系统对接落地配业务，其它场景请不要使用
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 品牌编码
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 扩展参数
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 月结卡号
	CustomerCode string `json:"customer_code,omitempty" xml:"customer_code,omitempty"`
	// 预约上门收件时间
	DoorPickUpTime string `json:"door_pick_up_time,omitempty" xml:"door_pick_up_time,omitempty"`
	// 预约上门截止时间
	DoorPickUpEndTime string `json:"door_pick_up_end_time,omitempty" xml:"door_pick_up_end_time,omitempty"`
	// 寄件网点编码
	ShippingBranchCode string `json:"shipping_branch_code,omitempty" xml:"shipping_branch_code,omitempty"`
	// 发货人信息
	Sender *UserInfoDto `json:"sender,omitempty" xml:"sender,omitempty"`
	// 是否使用智分宝预分拣， 仓库WMS系统对接落地配业务，其它场景请不要使用
	DmsSorting bool `json:"dms_sorting,omitempty" xml:"dms_sorting,omitempty"`
	// 订单上是否带3PLtiming属性, 该属性需要严格与订单上属性保持一致，如果不确定，请使用默认false。
	ThreePlTiming bool `json:"three_pl_timing,omitempty" xml:"three_pl_timing,omitempty"`
	// 设定取号返回的云打印报文是否加密
	NeedEncrypt bool `json:"need_encrypt,omitempty" xml:"need_encrypt,omitempty"`
	// 快递公司支持一票多件，快运公司子母件请勿使用该参数
	MultiPackagesShipment bool `json:"multi_packages_shipment,omitempty" xml:"multi_packages_shipment,omitempty"`
	// 是否预约上门
	CallDoorPickUp bool `json:"call_door_pick_up,omitempty" xml:"call_door_pick_up,omitempty"`
}
