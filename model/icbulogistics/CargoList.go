package icbulogistics

// CargoList 
type CargoList struct {

    // 单位
    
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    

    // 海关编码
    
    Hscode   string `json:"hscode,omitempty" xml:"hscode,omitempty"`
    

    // 货物数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 1
    
    DeclarationValue   string `json:"declaration_value,omitempty" xml:"declaration_value,omitempty"`
    

    // 货物单价
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 货物中文名
    
    NameCn   string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
    

    // 1
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 货物英文名
    
    NameEn   string `json:"name_en,omitempty" xml:"name_en,omitempty"`
    

    // 商品特性列表对象
    
    ProductType   []ProductType `json:"product_type,omitempty" xml:"product_type,omitempty"`
    

    // 材质
    
    Material   string `json:"material,omitempty" xml:"material,omitempty"`
    

    // 用途
    
    Purpose   string `json:"purpose,omitempty" xml:"purpose,omitempty"`
    

}
