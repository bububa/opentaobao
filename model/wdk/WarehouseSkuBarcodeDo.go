package wdk

// WarehouseSkuBarcodeDo 
type WarehouseSkuBarcodeDo struct {

    // 条码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 高
    
    Height   string `json:"height,omitempty" xml:"height,omitempty"`
    

    // 长
    
    Length   string `json:"length,omitempty" xml:"length,omitempty"`
    

    // 是否主条码
    
    MainFlag   bool `json:"main_flag,omitempty" xml:"main_flag,omitempty"`
    

    // 规格
    
    SpuSpec   string `json:"spu_spec,omitempty" xml:"spu_spec,omitempty"`
    

    // 单位
    
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    

    // 重量，单位g
    
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    

    // 宽
    
    Width   string `json:"width,omitempty" xml:"width,omitempty"`
    

}
