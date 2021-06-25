package jst

// CancelOrderRequest 
type CancelOrderRequest struct {

    // 规格，目前只有一个规格“HOV”，表示覆盖华为/OPPO/vivo
    Spec   string `json:"spec,omitempty"`

}
