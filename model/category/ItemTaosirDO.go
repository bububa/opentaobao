package category

// ItemTaosirDO 
type ItemTaosirDO struct {

    // 卖家可选单位List<单位id，单位名>
    
    StdUnitList   []Feature `json:"std_unit_list,omitempty" xml:"std_unit_list,omitempty"`
    

    // 表达式元素list
    
    ExprElList   []ItemTaoSirElDO `json:"expr_el_list,omitempty" xml:"expr_el_list,omitempty"`
    

    // 时间类型：0表示非时间，1表示时间点，2表示时间范围
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 数值小数点精度
    
    Precision   int64 `json:"precision,omitempty" xml:"precision,omitempty"`
    

}
