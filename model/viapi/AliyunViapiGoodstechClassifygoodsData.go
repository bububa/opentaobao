package viapi

import (
	"sync"
)

// AliyunViapiGoodstechClassifygoodsData 结构体
type AliyunViapiGoodstechClassifygoodsData struct {
	// 类目预测列表
	CategoryList []Category `json:"category_list,omitempty" xml:"category_list>category,omitempty"`
}

var poolAliyunViapiGoodstechClassifygoodsData = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechClassifygoodsData)
	},
}

// GetAliyunViapiGoodstechClassifygoodsData() 从对象池中获取AliyunViapiGoodstechClassifygoodsData
func GetAliyunViapiGoodstechClassifygoodsData() *AliyunViapiGoodstechClassifygoodsData {
	return poolAliyunViapiGoodstechClassifygoodsData.Get().(*AliyunViapiGoodstechClassifygoodsData)
}

// ReleaseAliyunViapiGoodstechClassifygoodsData 释放AliyunViapiGoodstechClassifygoodsData
func ReleaseAliyunViapiGoodstechClassifygoodsData(v *AliyunViapiGoodstechClassifygoodsData) {
	v.CategoryList = v.CategoryList[:0]
	poolAliyunViapiGoodstechClassifygoodsData.Put(v)
}
