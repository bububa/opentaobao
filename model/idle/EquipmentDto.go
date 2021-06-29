package idle

// EquipmentDto 
type EquipmentDto struct {

    // 标配名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 标配值，多为数量
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

}
