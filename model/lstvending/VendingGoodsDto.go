package lstvending

import (
	"github.com/bububa/opentaobao/model"
)

// VendingGoodsDto 结构体
type VendingGoodsDto struct {
	// 商品图片内容字节数组
	ImgData []*model.File `json:"img_data,omitempty" xml:"img_data>*model.File,omitempty"`
	// 建议摆放的货架层数
	ShelfNoList []int64 `json:"shelf_no_list,omitempty" xml:"shelf_no_list>int64,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 121212133
	ExternalId string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	// 计量单位，如：个、件、包
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品分类
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 厂商设备唯一编码
	EquipmentCode string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
	// 商品图片ID
	ImgPathId string `json:"img_path_id,omitempty" xml:"img_path_id,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 商品图片访问地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 状态：1上架，2下架，3删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品建议零售价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 商品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
