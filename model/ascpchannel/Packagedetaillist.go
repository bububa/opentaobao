package ascpchannel

// Packagedetaillist 
type Packagedetaillist struct {
    // 包件体积(m3)
    PackageVolume   string `json:"package_volume,omitempty" xml:"package_volume,omitempty"`
    // 包件重量(kg)
    PackageWeight   string `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
    // 包件名称
    PackageName   string `json:"package_name,omitempty" xml:"package_name,omitempty"`
    // 扩展字段JSON串
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
}
