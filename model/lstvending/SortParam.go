package lstvending

// SortParam 
type SortParam struct {

    // 排序字段名称：gmt_create创建时间、gmt_modified修改时间、id主键
    
    SortFieldName   string `json:"sort_field_name,omitempty" xml:"sort_field_name,omitempty"`
    

    // 排序方式：asc、desc
    
    SortOrder   string `json:"sort_order,omitempty" xml:"sort_order,omitempty"`
    

}
