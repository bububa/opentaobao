package alidoc

// DrugStoreAddTopRequest 
type DrugStoreAddTopRequest struct {

    // 药店地址
    
    DrugStoreAddress   string `json:"drug_store_address,omitempty" xml:"drug_store_address,omitempty"`
    

    // 维度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 电话
    
    DrugStrorePhone   string `json:"drug_strore_phone,omitempty" xml:"drug_strore_phone,omitempty"`
    

    // 医保标签
    
    MedicareLabel   int64 `json:"medicare_label,omitempty" xml:"medicare_label,omitempty"`
    

    // 药店编码
    
    DrugStoreCode   string `json:"drug_store_code,omitempty" xml:"drug_store_code,omitempty"`
    

    // 经度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

    // 药店名称
    
    DrugStoreName   string `json:"drug_store_name,omitempty" xml:"drug_store_name,omitempty"`
    

}
