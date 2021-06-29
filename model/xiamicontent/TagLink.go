package xiamicontent

// TagLink 
type TagLink struct {
    // tag直属类目
    Parent   *TagCatLink `json:"parent,omitempty" xml:"parent,omitempty"`
    // tag code(唯一标识一个tag)
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // tag名称
    NameCn   string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
    // tag描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 父类目id
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`
    // tag英文名
    NameEn   string `json:"name_en,omitempty" xml:"name_en,omitempty"`
    // tag id(唯一标识一个tag)
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
