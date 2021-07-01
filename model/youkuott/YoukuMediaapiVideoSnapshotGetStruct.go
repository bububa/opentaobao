package youkuott

// YoukuMediaapiVideoSnapshotGetStruct 结构体
type YoukuMediaapiVideoSnapshotGetStruct struct {
	// 图片域名
	DomainName string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// 图片url列表
	ThumbIdList []string `json:"thumb_id_list,omitempty" xml:"thumb_id_list>string,omitempty"`
	// 毫秒
	Sectiontime int64 `json:"sectiontime,omitempty" xml:"sectiontime,omitempty"`
}
