package omniorder

// TaobaoomniorderstorecollectconfiggetResult 结构体
type TaobaoomniorderstorecollectconfiggetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *StoreCollectConfig `json:"data,omitempty" xml:"data,omitempty"`
}
