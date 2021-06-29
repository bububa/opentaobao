package idle

// MediaDto 
type MediaDto struct {

    // 商品主图列表
    
    Images   []ImageInfoDto `json:"images,omitempty" xml:"images,omitempty"`
    

    // 商品详情图片列表
    
    PropImages   []ImageInfoDto `json:"prop_images,omitempty" xml:"prop_images,omitempty"`
    

}
