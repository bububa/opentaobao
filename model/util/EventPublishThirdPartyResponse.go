package util

// EventPublishThirdPartyResponse 结构体
type EventPublishThirdPartyResponse struct {
	// 发布事件结果列表
	EntryList []EventPublishThirdPartyResultEntry `json:"entry_list,omitempty" xml:"entry_list>event_publish_third_party_result_entry,omitempty"`
}
