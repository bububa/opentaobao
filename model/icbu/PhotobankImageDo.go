package icbu

// PhotobankImageDo 结构体
type PhotobankImageDo struct {
	// 图片url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 文件名字
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 归属人
	OwnerMemberDisplayName string `json:"owner_member_display_name,omitempty" xml:"owner_member_display_name,omitempty"`
	// 展示名字
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
	// 图片引用数量
	ReferenceCount int64 `json:"reference_count,omitempty" xml:"reference_count,omitempty"`
	// 分组id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
}
