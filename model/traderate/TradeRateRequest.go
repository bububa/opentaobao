package traderate

// TradeRateRequest 结构体
type TradeRateRequest struct {
	//
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	//
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	//
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
