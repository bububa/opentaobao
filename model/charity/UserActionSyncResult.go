package charity

// UserActionSyncResult 结构体
type UserActionSyncResult struct {
	// 唯一的事件ID
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 本次同步获得的公益时
	CharityHours string `json:"charity_hours,omitempty" xml:"charity_hours,omitempty"`
}
