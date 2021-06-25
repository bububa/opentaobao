package product

// ResourceDataRecord 
type ResourceDataRecord struct {

    // 映射id
    MappingId   int64 `json:"mapping_id,omitempty"`

    // 项目id
    ProductId   int64 `json:"product_id,omitempty"`

    // 次数
    Num   int64 `json:"num,omitempty"`

    // 资源类型
    Type   int64 `json:"type,omitempty"`

    // 资源id
    Id   int64 `json:"id,omitempty"`

    // 扩展属性，包括城市名称，外部编码等
    ValueMap   string `json:"value_map,omitempty"`

}
