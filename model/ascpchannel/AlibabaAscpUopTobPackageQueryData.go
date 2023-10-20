package ascpchannel

import (
	"sync"
)

// AlibabaAscpUopTobPackageQueryData 结构体
type AlibabaAscpUopTobPackageQueryData struct {
	// 包裹明细列表
	PackageItemDtos []Packageitemdtos `json:"package_item_dtos,omitempty" xml:"package_item_dtos>packageitemdtos,omitempty"`
	// 包裹信息
	PackageDto *Packagedto `json:"package_dto,omitempty" xml:"package_dto,omitempty"`
}

var poolAlibabaAscpUopTobPackageQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopTobPackageQueryData)
	},
}

// GetAlibabaAscpUopTobPackageQueryData() 从对象池中获取AlibabaAscpUopTobPackageQueryData
func GetAlibabaAscpUopTobPackageQueryData() *AlibabaAscpUopTobPackageQueryData {
	return poolAlibabaAscpUopTobPackageQueryData.Get().(*AlibabaAscpUopTobPackageQueryData)
}

// ReleaseAlibabaAscpUopTobPackageQueryData 释放AlibabaAscpUopTobPackageQueryData
func ReleaseAlibabaAscpUopTobPackageQueryData(v *AlibabaAscpUopTobPackageQueryData) {
	v.PackageItemDtos = v.PackageItemDtos[:0]
	v.PackageDto = nil
	poolAlibabaAscpUopTobPackageQueryData.Put(v)
}
