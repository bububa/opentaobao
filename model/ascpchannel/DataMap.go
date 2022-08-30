package ascpchannel

// DataMap 结构体
type DataMap struct {
	// 幂等原因
	IdempotentReason string `json:"idempotent_reason,omitempty" xml:"idempotent_reason,omitempty"`
}
