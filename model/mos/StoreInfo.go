package mos

// StoreInfo 结构体
type StoreInfo struct {
	// 图片
	StorePic string `json:"store_pic,omitempty" xml:"store_pic,omitempty"`
	// 专柜名字
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 专柜ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
