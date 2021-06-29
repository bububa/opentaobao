package product

// ResourceData 
type ResourceData struct {
    // 资源列表
    Datas   []ResourceDataRecord `json:"datas,omitempty" xml:"datas>resource_data_record,omitempty"`
    // 资源名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 资源类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
