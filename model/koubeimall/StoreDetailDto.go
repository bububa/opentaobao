package koubeimall

// StoreDetailDTO 
type StoreDetailDTO struct {
    // 门店相册
    StoreAlbum   *StoreAlbumDTO `json:"store_album,omitempty" xml:"store_album,omitempty"`
    // 门店基础信息
    StoreDto   *StoreDTO `json:"store_dto,omitempty" xml:"store_dto,omitempty"`
    // 服务信息
    ServiceInfo   *ServiceInfoDTO `json:"service_info,omitempty" xml:"service_info,omitempty"`
}
