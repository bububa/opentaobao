package hotel

// ShotelPropertiesVo 
type ShotelPropertiesVo struct {

    // 条目id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 排序
    
    OrderNum   int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
    

    // 所属shid
    
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    

    // 所属srid，如果是酒店上属性，则该字段无值
    
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    

    // 二级分类
    
    SubType   string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
    

    // 一级分类，10-预订须知
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 二级服务设施字典code,10001--入店时间, 10002--离店时间, 10003--酒店政策, 10004--加床收费, 10005--额外收费, 10006--其他说明, 10007--酒店位于山顶, 10008--需坐船前往, 10009--位于景区内
    
    TypeId   string `json:"type_id,omitempty" xml:"type_id,omitempty"`
    

    // 实际内容
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

}
