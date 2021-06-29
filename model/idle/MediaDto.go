package idle

// MediaDTO 
type MediaDTO struct {
    // 商品主图列表
    Images   []ImageInfoDTO `json:"images,omitempty" xml:"images>image_info_dto,omitempty"`
    // 商品详情图片列表
    PropImages   []ImageInfoDTO `json:"prop_images,omitempty" xml:"prop_images>image_info_dto,omitempty"`
}
