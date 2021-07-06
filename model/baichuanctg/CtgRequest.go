package baichuanctg

// CtgRequest 结构体
type CtgRequest struct {
	// delivery_id
	DeliveryId string `json:"delivery_id,omitempty" xml:"delivery_id,omitempty"`
	// res_id
	ResId string `json:"res_id,omitempty" xml:"res_id,omitempty"`
	// app_key
	BusinessAppKey string `json:"business_app_key,omitempty" xml:"business_app_key,omitempty"`
	// date
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// page_size
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// current_page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
