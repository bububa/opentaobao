package openmall

// TopParseAddressVo 结构体
type TopParseAddressVo struct {
	// 地址解析结构
	Entries []TopParseAddressEntryVo `json:"entries,omitempty" xml:"entries>top_parse_address_entry_vo,omitempty"`
}
