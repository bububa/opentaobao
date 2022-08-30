package ascp

// TmsOrderUpdateRequest 结构体
type TmsOrderUpdateRequest struct {
	// 货品信息
	ItemEditDtoList []OrderItemEdit `json:"item_edit_dto_list,omitempty" xml:"item_edit_dto_list>order_item_edit,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 服务商code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 底单图片（多个url以逗号分开）
	BillPic string `json:"bill_pic,omitempty" xml:"bill_pic,omitempty"`
	// 行业标(家装固定值，home_improvement)
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 外部单号
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 操作人
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 服务类型 (0: “干支装一体”1: “不安装配送到家”2:“不安装配送到楼下”3: “自提”)
	TmsServiceType int64 `json:"tms_service_type,omitempty" xml:"tms_service_type,omitempty"`
	// 拼几单
	GatherNum int64 `json:"gather_num,omitempty" xml:"gather_num,omitempty"`
	// 发货人信息
	SenderEditDto *OrderSenderEdit `json:"sender_edit_dto,omitempty" xml:"sender_edit_dto,omitempty"`
	// 收货人信息
	ReceiverEditDto *OrderReceiverEdit `json:"receiver_edit_dto,omitempty" xml:"receiver_edit_dto,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 是否拼单
	Gather bool `json:"gather,omitempty" xml:"gather,omitempty"`
}
