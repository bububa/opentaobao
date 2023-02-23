package alihealth2

// ValidityPeriodSyncReqDto 结构体
type ValidityPeriodSyncReqDto struct {
	// 请求明细
	Items []ValidityPeriodItem `json:"items,omitempty" xml:"items>validity_period_item,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
