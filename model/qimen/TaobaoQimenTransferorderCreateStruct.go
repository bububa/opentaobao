package qimen

// TaobaoQimenTransferorderCreateStruct 结构体
type TaobaoQimenTransferorderCreateStruct struct {
	// 项目集
	TransferItems []TransferItems `json:"transferItems,omitempty" xml:"transferItems>transfer_items,omitempty"`
	// 外部ERP订单号,HZ1234,string(50),必填,
	ErpOrderCode string `json:"erpOrderCode,omitempty" xml:"erpOrderCode,omitempty"`
	// 出库仓编码,Item1234,string(50),必填,
	FromStoreCode string `json:"fromStoreCode,omitempty" xml:"fromStoreCode,omitempty"`
	// 入库仓编码,HZ1234,string(50),必填,
	ToStoreCode string `json:"toStoreCode,omitempty" xml:"toStoreCode,omitempty"`
	// 期望调拨开始时间,Item1234,string(50),,
	ExpectStartTime string `json:"expectStartTime,omitempty" xml:"expectStartTime,omitempty"`
	// 扩展属性,HZ1234,string(500),,
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 111
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 响应结果:success|failure,success,string(10),必填,
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,0,string(50),,
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,invalid appkey,string(100),,
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调拨单号,0,string(50),,
	TransferOrderCode string `json:"transferOrderCode,omitempty" xml:"transferOrderCode,omitempty"`
	// 预计出库时间,0,string(50),,
	ExpectOutStoreTime string `json:"expectOutStoreTime,omitempty" xml:"expectOutStoreTime,omitempty"`
	// 预计入库时间,0,string(50),,
	ExpectInStoreTime string `json:"expectInStoreTime,omitempty" xml:"expectInStoreTime,omitempty"`
	// 调拨单信息
	TransferExecuteInfo *TaobaoQimenTransferorderCreateStruct `json:"transferExecuteInfo,omitempty" xml:"transferExecuteInfo,omitempty"`
}
