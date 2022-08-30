package damaiticklet

// ScriptTagThirdParam 结构体
type ScriptTagThirdParam struct {
	// 剧本名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 剧本类型,核心题材1   剧本题材 2  时代背景 3   流派/类型 4  剧本类型 5
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 剧本标签id
	OutTagId int64 `json:"out_tag_id,omitempty" xml:"out_tag_id,omitempty"`
}
