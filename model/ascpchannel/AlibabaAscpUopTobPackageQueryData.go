package ascpchannel

// AlibabaAscpUopTobPackageQueryData 结构体
type AlibabaAscpUopTobPackageQueryData struct {
	// 包裹明细列表
	PackageItemDtos []Packageitemdtos `json:"package_item_dtos,omitempty" xml:"package_item_dtos>packageitemdtos,omitempty"`
	// 包裹信息
	PackageDto *Packagedto `json:"package_dto,omitempty" xml:"package_dto,omitempty"`
}
