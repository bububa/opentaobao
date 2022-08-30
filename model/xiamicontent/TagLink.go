package xiamicontent

// TagLink 结构体
type TagLink struct {
	// 标签code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 标签名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// tag描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// tag英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 标签id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 父标签id
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// tag直属类目
	Parent *TagCatLink `json:"parent,omitempty" xml:"parent,omitempty"`
}
