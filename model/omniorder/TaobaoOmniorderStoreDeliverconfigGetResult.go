package omniorder

// TaobaoOmniorderStoreDeliverconfigGetResult 结构体
type TaobaoOmniorderStoreDeliverconfigGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *StoreDeliverConfig `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
