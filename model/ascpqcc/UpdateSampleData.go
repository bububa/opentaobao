package ascpqcc

import (
	"sync"
)

// UpdateSampleData 结构体
type UpdateSampleData struct {
	// 厂商商品尺寸
	VendorSize string `json:"vendor_size,omitempty" xml:"vendor_size,omitempty"`
	// 业务系统商品编号
	BizItemId string `json:"biz_item_id,omitempty" xml:"biz_item_id,omitempty"`
}

var poolUpdateSampleData = sync.Pool{
	New: func() any {
		return new(UpdateSampleData)
	},
}

// GetUpdateSampleData() 从对象池中获取UpdateSampleData
func GetUpdateSampleData() *UpdateSampleData {
	return poolUpdateSampleData.Get().(*UpdateSampleData)
}

// ReleaseUpdateSampleData 释放UpdateSampleData
func ReleaseUpdateSampleData(v *UpdateSampleData) {
	v.VendorSize = ""
	v.BizItemId = ""
	poolUpdateSampleData.Put(v)
}
