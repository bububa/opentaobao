package ascp

// TmsOrderCreateRequest 结构体
type TmsOrderCreateRequest struct {
	// 货品信息
	ItemCreateDtoList []OrderReceiverCreate `json:"item_create_dto_list,omitempty" xml:"item_create_dto_list>order_receiver_create,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 交易单号
	TradeCode string `json:"trade_code,omitempty" xml:"trade_code,omitempty"`
	// 外部单号
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 商家单号
	SellerOrderCode string `json:"seller_order_code,omitempty" xml:"seller_order_code,omitempty"`
	// 服务商code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 创建者
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 底单图片（多个url以逗号分开）
	BillPic string `json:"bill_pic,omitempty" xml:"bill_pic,omitempty"`
	// 业务请求ID,用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 行业标(家装固定值，home_improvement)
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 服务类型 (0: “干支装一体”1: “不安装配送到家”2:“不安装配送到楼下”3: “自提”)
	TmsServiceType int64 `json:"tms_service_type,omitempty" xml:"tms_service_type,omitempty"`
	// 拼几单
	GatherNum int64 `json:"gather_num,omitempty" xml:"gather_num,omitempty"`
	// 业务请求时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 收货分拨信息
	RdcCreateDto *OrderRdcCreate `json:"rdc_create_dto,omitempty" xml:"rdc_create_dto,omitempty"`
	// 发货人信息
	SenderCreateDto *OrderSenderCreate `json:"sender_create_dto,omitempty" xml:"sender_create_dto,omitempty"`
	// 收货人信息
	ReceiverCreateDto *OrderSenderCreate `json:"receiver_create_dto,omitempty" xml:"receiver_create_dto,omitempty"`
	// 是否预开单
	Advance bool `json:"advance,omitempty" xml:"advance,omitempty"`
	// 是否拼单
	Gather bool `json:"gather,omitempty" xml:"gather,omitempty"`
}
