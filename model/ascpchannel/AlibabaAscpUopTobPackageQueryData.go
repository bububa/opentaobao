package ascpchannel

// AlibabaAscpUopTobPackageQueryData 
type AlibabaAscpUopTobPackageQueryData struct {

    // 包裹信息
    
    PackageDto   *Packagedto `json:"package_dto,omitempty" xml:"package_dto,omitempty"`
    

    // 包裹明细列表
    
    PackageItemDtos   []Packageitemdtos `json:"package_item_dtos,omitempty" xml:"package_item_dtos,omitempty"`
    

}
