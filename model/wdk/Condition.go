package wdk

// Condition 
type Condition struct {

    // 满元金额，单位分
    Amount   int64 `json:"amount,omitempty"`

    // 是否满件
    CountAt   bool `json:"count_at,omitempty"`

    // 是否指定件数
    CountBegin   bool `json:"count_begin,omitempty"`

    // 是否第N件
    Appoint   bool `json:"appoint,omitempty"`

    // 满件门槛
    Count   int64 `json:"count,omitempty"`

    // 是否满元
    AmountAt   bool `json:"amount_at,omitempty"`

}
