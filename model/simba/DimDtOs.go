package simba

// DimDtOs 
type DimDtOs struct {
    // tagList
    TagList   []TagOptions `json:"tag_list,omitempty" xml:"tag_list>tag_options,omitempty"`
    // 维度名称,如性别,年龄
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 维度id,如性别年龄的id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
