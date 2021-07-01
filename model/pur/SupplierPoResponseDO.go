package pur

// SupplierPoResponseDo 结构体
type SupplierPoResponseDo struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 反馈状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// po头编号
	PoNo string `json:"po_no,omitempty" xml:"po_no,omitempty"`
	// po行编号
	PoLineNum string `json:"po_line_num,omitempty" xml:"po_line_num,omitempty"`
	// 购买方式
	ProcurementMethod string `json:"procurement_method,omitempty" xml:"procurement_method,omitempty"`
	// 延误原因
	DelayReason string `json:"delay_reason,omitempty" xml:"delay_reason,omitempty"`
	// 发货金额
	DeliveryAmount string `json:"delivery_amount,omitempty" xml:"delivery_amount,omitempty"`
	// 产能计划代码
	CapacityRecordNo string `json:"capacity_record_no,omitempty" xml:"capacity_record_no,omitempty"`
	// 反馈类型
	ResponseType string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// 预计到货日期
	EstimatedArrivalDate string `json:"estimated_arrival_date,omitempty" xml:"estimated_arrival_date,omitempty"`
	// 生产周期
	ProductionPeriod string `json:"production_period,omitempty" xml:"production_period,omitempty"`
	// 采购组织
	DemanderPurchaseOrgCode string `json:"demander_purchase_org_code,omitempty" xml:"demander_purchase_org_code,omitempty"`
	// 发货数量
	DeliveryQty string `json:"delivery_qty,omitempty" xml:"delivery_qty,omitempty"`
	// 实际发货日期
	ActualArrivalDate string `json:"actual_arrival_date,omitempty" xml:"actual_arrival_date,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 项目代码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 供应商ou代码
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 发货批次号
	DeliveryBatchNo string `json:"delivery_batch_no,omitempty" xml:"delivery_batch_no,omitempty"`
}
