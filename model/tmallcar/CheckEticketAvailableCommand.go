package tmallcar

// CheckEticketAvailableCommand 结构体
type CheckEticketAvailableCommand struct {
	// 核销码
	EticketCode string `json:"eticket_code,omitempty" xml:"eticket_code,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
