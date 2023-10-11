package alicom

// SequceCallInfo 结构体
type SequceCallInfo struct {
	// ff
	SequenceCallEvents []SequenceCallEvent `json:"sequence_call_events,omitempty" xml:"sequence_call_events>sequence_call_event,omitempty"`
	// 顺振call的次数
	SequenceCallCount int64 `json:"sequence_call_count,omitempty" xml:"sequence_call_count,omitempty"`
}
