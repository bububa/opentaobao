package hotel

// HotelPackageVo 结构体
type HotelPackageVo struct {
	// detailPackageDesc
	DetailPackageDesc string `json:"detail_package_desc,omitempty" xml:"detail_package_desc,omitempty"`
	// detailPackageLogo
	DetailPackageLogo string `json:"detail_package_logo,omitempty" xml:"detail_package_logo,omitempty"`
	// howToPlay
	HowToPlay string `json:"how_to_play,omitempty" xml:"how_to_play,omitempty"`
	// note
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// packageDesc
	PackageDesc string `json:"package_desc,omitempty" xml:"package_desc,omitempty"`
	// packageId
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// packageName
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// packageType
	PackageType int64 `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// packageTypeName
	PackageTypeName string `json:"package_type_name,omitempty" xml:"package_type_name,omitempty"`
	// scenicAddress
	ScenicAddress string `json:"scenic_address,omitempty" xml:"scenic_address,omitempty"`
	// scenicCoverImg
	ScenicCoverImg string `json:"scenic_cover_img,omitempty" xml:"scenic_cover_img,omitempty"`
	// scenicName
	ScenicName string `json:"scenic_name,omitempty" xml:"scenic_name,omitempty"`
}
