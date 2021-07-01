package omniorder

// TaobaoOmniorderStoreCollectconfigGetResult 结构体
type TaobaoOmniorderStoreCollectconfigGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *StoreCollectConfig `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
