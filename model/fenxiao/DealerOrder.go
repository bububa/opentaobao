package fenxiao

// DealerOrder 结构体
type DealerOrder struct {
	// 经销采购单编号，API发货使用此字段
	DealerOrderId int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`
	// 供应商nick
	SupplierNick string `json:"supplier_nick,omitempty" xml:"supplier_nick,omitempty"`
	// 分销商nick
	ApplierNick string `json:"applier_nick,omitempty" xml:"applier_nick,omitempty"`
	// 供应商驳回申请的原因
	RefuseReasonSupplier string `json:"refuse_reason_supplier,omitempty" xml:"refuse_reason_supplier,omitempty"`
	// 分销商拒绝供应商修改的原因
	RefuseReasonApplier string `json:"refuse_reason_applier,omitempty" xml:"refuse_reason_applier,omitempty"`
	// 关闭原因
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 总采购数量
	QuantityCount int64 `json:"quantity_count,omitempty" xml:"quantity_count,omitempty"`
	// 采购总价(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	TotalPrice float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 申请时间
	AppliedTime string `json:"applied_time,omitempty" xml:"applied_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 供应商最后一次审核通过/修改/驳回的时间
	AuditTimeSupplier string `json:"audit_time_supplier,omitempty" xml:"audit_time_supplier,omitempty"`
	// 分销商最后一次确认/申请/拒绝的时间
	AuditTimeApplier string `json:"audit_time_applier,omitempty" xml:"audit_time_applier,omitempty"`
	// 已发货数量
	DeliveredQuantityCount int64 `json:"delivered_quantity_count,omitempty" xml:"delivered_quantity_count,omitempty"`
	// WAIT_FOR_SUPPLIER_AUDIT1：分销商提交申请，待供应商审核；SUPPLIER_REFUSE：供应商驳回申请，待分销商确认；WAIT_FOR_APPLIER_AUDIT：供应商修改后，待分销商确认；WAIT_FOR_SUPPLIER_AUDIT2：分销商拒绝修改，待供应商再审核；BOTH_AGREE_WAIT_PAY：审核通过下单成功，待分销商付款WAIT_FOR_SUPPLIER_DELIVER：付款成功，待供应商发货；WAIT_FOR_APPLIER_STORAGE：供应商发货，待分销商收货；TRADE_FINISHED：分销商收货，交易成功；TRADE_CLOSED：经销采购单关闭。
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 物流费用(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	LogisticsFee float64 `json:"logistics_fee,omitempty" xml:"logistics_fee,omitempty"`
	// 支付方式：ALIPAY_SURETY（支付宝担保交易）TRANSFER（线下转账）PREPAY（预存款）IMMEDIATELY（即时到账）
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 物流方式：SELF_PICKUP（自提）、LOGISTICS（物流发货)
	LogisticsType string `json:"logistics_type,omitempty" xml:"logistics_type,omitempty"`
	// 收货人信息
	Receiver *Receiver `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 产品明细
	DealerOrderDetails []DealerOrderDetail `json:"dealer_order_details,omitempty" xml:"dealer_order_details>dealer_order_detail,omitempty"`
	// 付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付宝交易号
	AlipayNo string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// 供应商备注。仅供应商可见。
	SupplierMemo string `json:"supplier_memo,omitempty" xml:"supplier_memo,omitempty"`
	// 供应商备注旗帜。1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色。仅供应商可见。
	SupplierMemoFlag int64 `json:"supplier_memo_flag,omitempty" xml:"supplier_memo_flag,omitempty"`
	// 折让费用(精确到2位小数;单位:元。如:200.07，表示:200元7分 )
	RebateFee float64 `json:"rebate_fee,omitempty" xml:"rebate_fee,omitempty"`
	// 属性信息列表，key-value形式。
	Features []Feature `json:"features,omitempty" xml:"features>feature,omitempty"`
	// 属性键
	DistMemo string `json:"dist_memo,omitempty" xml:"dist_memo,omitempty"`
}
