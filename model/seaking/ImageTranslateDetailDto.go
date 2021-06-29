package seaking

// ImageTranslateDetailDto 
type ImageTranslateDetailDto struct {

    // 目标语种
    
    TargetLang   string `json:"target_lang,omitempty" xml:"target_lang,omitempty"`
    

    // 源语种
    
    SourceLang   string `json:"source_lang,omitempty" xml:"source_lang,omitempty"`
    

    // 图片url
    
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    

    // 商品id
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 商品所在平台（ae/lazada)）
    
    Platform   string `json:"platform,omitempty" xml:"platform,omitempty"`
    

    // 子任务编号，从1开始，必须连续
    
    Idx   int64 `json:"idx,omitempty" xml:"idx,omitempty"`
    

}
