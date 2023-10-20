package ascpchannel

// AlibabaascpuoptobpackagequeryData 结构体
type AlibabaascpuoptobpackagequeryData struct {
	// 包裹明细列表
	PackageItemDtos []Packageitemdtos `json:"package_item_dtos,omitempty" xml:"package_item_dtos>packageitemdtos,omitempty"`
	// 包裹信息
	PackageDto *Packagedto `json:"package_dto,omitempty" xml:"package_dto,omitempty"`
}
