package omniorder

// StoreAcceptedResult 结构体
type StoreAcceptedResult struct {
	// 主订单Id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 店铺类型, 门店或者电商仓
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 店铺Id, 可能是门店或者电商仓
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 子订单Id
	SubOid int64 `json:"sub_oid,omitempty" xml:"sub_oid,omitempty"`
	// 异常描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0表示无系统异常
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
