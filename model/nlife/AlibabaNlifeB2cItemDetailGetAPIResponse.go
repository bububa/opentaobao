package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cItemDetailGetAPIResponse b2c码详情查询 API返回值
// alibaba.nlife.b2c.item.detail.get
//
// 根据零售+平台生成的唯一码获取对应详情
type AlibabaNlifeB2cItemDetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cItemDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cItemDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cItemDetailGetAPIResponseModel).Reset()
}

// AlibabaNlifeB2cItemDetailGetAPIResponseModel is b2c码详情查询 成功返回结果
type AlibabaNlifeB2cItemDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_item_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品图片链接(线下商品无)
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商品出售价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品原价
	OrgPrice string `json:"org_price,omitempty" xml:"org_price,omitempty"`
	// 商品在天猫上的详情页链接(线下商品无
	DetailUrl string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
	// 商品ItemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品SkuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku级别
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 零售+平台生成的商品唯一码或条码
	UniqueCode string `json:"unique_code,omitempty" xml:"unique_code,omitempty"`
	// 入驻天猫的品牌ID(线下商品无
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 入驻天猫的品牌名称(线下商品无)
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 商品级别
	GoodsNo string `json:"goods_no,omitempty" xml:"goods_no,omitempty"`
	// 结算码
	SettleCode string `json:"settle_code,omitempty" xml:"settle_code,omitempty"`
	// 销售属性
	Property string `json:"property,omitempty" xml:"property,omitempty"`
	// 叶子类目ID
	CatId string `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 叶子类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 当前sku库存
	CurrentInventory string `json:"current_inventory,omitempty" xml:"current_inventory,omitempty"`
	// 结算码是否可变
	CodeChangeable bool `json:"code_changeable,omitempty" xml:"code_changeable,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cItemDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PicUrl = ""
	m.Name = ""
	m.Price = ""
	m.OrgPrice = ""
	m.DetailUrl = ""
	m.ItemId = ""
	m.SkuId = ""
	m.BarCode = ""
	m.UniqueCode = ""
	m.BrandId = ""
	m.BrandName = ""
	m.SupplierId = ""
	m.SupplierName = ""
	m.GoodsNo = ""
	m.SettleCode = ""
	m.Property = ""
	m.CatId = ""
	m.CatName = ""
	m.CurrentInventory = ""
	m.CodeChangeable = false
}

var poolAlibabaNlifeB2cItemDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cItemDetailGetAPIResponse)
	},
}

// GetAlibabaNlifeB2cItemDetailGetAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cItemDetailGetAPIResponse
func GetAlibabaNlifeB2cItemDetailGetAPIResponse() *AlibabaNlifeB2cItemDetailGetAPIResponse {
	return poolAlibabaNlifeB2cItemDetailGetAPIResponse.Get().(*AlibabaNlifeB2cItemDetailGetAPIResponse)
}

// ReleaseAlibabaNlifeB2cItemDetailGetAPIResponse 将 AlibabaNlifeB2cItemDetailGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cItemDetailGetAPIResponse(v *AlibabaNlifeB2cItemDetailGetAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cItemDetailGetAPIResponse.Put(v)
}
