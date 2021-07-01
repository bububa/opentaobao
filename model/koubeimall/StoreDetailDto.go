package koubeimall

// StoreDetailDto 
type StoreDetailDto struct {
    // 门店相册
    StoreAlbum   *StoreAlbumDto `json:"store_album,omitempty" xml:"store_album,omitempty"`
    // 门店基础信息
    StoreDto   *StoreDto `json:"store_dto,omitempty" xml:"store_dto,omitempty"`
    // 服务信息
    ServiceInfo   *ServiceInfoDto `json:"service_info,omitempty" xml:"service_info,omitempty"`
}
