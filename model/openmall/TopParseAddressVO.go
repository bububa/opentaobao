package openmall

// TopParseAddressVO 结构体
type TopParseAddressVO struct {
	// 地址解析结构
	Entries []TopParseAddressEntryVO `json:"entries,omitempty" xml:"entries>top_parse_address_entry_vo,omitempty"`
}
