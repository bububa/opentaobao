package hotel

// LikeTargetCount 
type LikeTargetCount struct {
    // count
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // targetId
    TargetId   int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
    // voted
    Voted   bool `json:"voted,omitempty" xml:"voted,omitempty"`
}
