package cainiaohandover

// OpenOrderParam 结构体
type OpenOrderParam struct {
	// 包裹参数
	PackageParams []OpenPackageParam `json:"package_params,omitempty" xml:"package_params>open_package_param,omitempty"`
	// 交易单参数
	TradeOrderParam *OpenTradeOrderParam `json:"trade_order_param,omitempty" xml:"trade_order_param,omitempty"`
	// 解决方案参数
	SolutionParam *OpenSolutionParam `json:"solution_param,omitempty" xml:"solution_param,omitempty"`
	// 商家信息参数
	SellerInfoParam *OpenSellerInfoParam `json:"seller_info_param,omitempty" xml:"seller_info_param,omitempty"`
	// 发件人信息
	SenderParam *OpenSenderParam `json:"sender_param,omitempty" xml:"sender_param,omitempty"`
	// 退货联系人信息
	ReturnerParam *OpenReturnerParam `json:"returner_param,omitempty" xml:"returner_param,omitempty"`
	// 收件人信息
	ReceiverParam *ReceiverParam `json:"receiver_param,omitempty" xml:"receiver_param,omitempty"`
	// 揽收信息参数
	PickupInfoParam *OpenPickupInfoParam `json:"pickup_info_param,omitempty" xml:"pickup_info_param,omitempty"`
}
