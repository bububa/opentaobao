package moscm

import (
	"sync"
)

// Spudto 结构体
type Spudto struct {
	// 商品图片集合
	ProductImgs []ProductImgDto `json:"product_imgs,omitempty" xml:"product_imgs>product_img_dto,omitempty"`
	// 属性
	Props []PropertyDto `json:"props,omitempty" xml:"props>property_dto,omitempty"`
	// 条码信息
	BarcodeStr string `json:"barcode_str,omitempty" xml:"barcode_str,omitempty"`
	// 品牌唯一标识
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 类目唯一标识
	Cid string `json:"cid,omitempty" xml:"cid,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 唯一标识
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 透明素材图
	Material string `json:"material,omitempty" xml:"material,omitempty"`
	// m站产品描述
	Mdesc string `json:"mdesc,omitempty" xml:"mdesc,omitempty"`
	// 修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// pc端产品描述
	PcDesc string `json:"pc_desc,omitempty" xml:"pc_desc,omitempty"`
	// 图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 吊牌价,单位:元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 产品唯一标识
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品卖点描述，长度限制在20个汉字
	SellPt string `json:"sell_pt,omitempty" xml:"sell_pt,omitempty"`
	// 款号
	StyleNo string `json:"style_no,omitempty" xml:"style_no,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 扩展属性
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 产品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品级别:1.天猫，2.线上，3.单品，4.原始
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 产品评分次数
	RateNum int64 `json:"rate_num,omitempty" xml:"rate_num,omitempty"`
	// 产品状态：删除(-1),正常(1)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否新品,默认true
	IsNew bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
}

var poolSpudto = sync.Pool{
	New: func() any {
		return new(Spudto)
	},
}

// GetSpudto() 从对象池中获取Spudto
func GetSpudto() *Spudto {
	return poolSpudto.Get().(*Spudto)
}

// ReleaseSpudto 释放Spudto
func ReleaseSpudto(v *Spudto) {
	v.ProductImgs = v.ProductImgs[:0]
	v.Props = v.Props[:0]
	v.BarcodeStr = ""
	v.BrandId = ""
	v.BrandName = ""
	v.CatName = ""
	v.Cid = ""
	v.Created = ""
	v.Id = ""
	v.Material = ""
	v.Mdesc = ""
	v.Modified = ""
	v.PcDesc = ""
	v.PicUrl = ""
	v.Price = ""
	v.ProductId = ""
	v.SellPt = ""
	v.StyleNo = ""
	v.SubTitle = ""
	v.Tags = ""
	v.Title = ""
	v.Level = 0
	v.RateNum = 0
	v.Status = 0
	v.IsNew = false
	poolSpudto.Put(v)
}
