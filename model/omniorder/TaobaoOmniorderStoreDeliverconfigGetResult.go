package omniorder

// TaobaoomniorderstoredeliverconfiggetResult 结构体
type TaobaoomniorderstoredeliverconfiggetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *StoreDeliverConfig `json:"data,omitempty" xml:"data,omitempty"`
}
