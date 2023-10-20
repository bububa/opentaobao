package omniorder

// TaobaoOmniorderStoreSdtstatusResult 结构体
type TaobaoOmniorderStoreSdtstatusResult struct {
	// 异常信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 异常码 0  正常，否则异常
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// data
	Data *SdtStatusResponse `json:"data,omitempty" xml:"data,omitempty"`
}
