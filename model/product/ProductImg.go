package product

import (
	"sync"
)

// ProductImg 结构体
type ProductImg struct {
	// 图片地址.(绝对地址,格式:http://host/image_path)
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 添加时间.格式:yyyy-mm-dd hh:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 修改时间.格式:yyyy-mm-dd hh:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 产品图片ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 图片所属产品的ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}

var poolProductImg = sync.Pool{
	New: func() any {
		return new(ProductImg)
	},
}

// GetProductImg() 从对象池中获取ProductImg
func GetProductImg() *ProductImg {
	return poolProductImg.Get().(*ProductImg)
}

// ReleaseProductImg 释放ProductImg
func ReleaseProductImg(v *ProductImg) {
	v.Url = ""
	v.Created = ""
	v.Modified = ""
	v.Id = 0
	v.ProductId = 0
	v.Position = 0
	poolProductImg.Put(v)
}
