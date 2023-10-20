package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemSkuDeleteAPIResponse 酒店团购套餐商品SKU删除 API返回值
// alitrip.tuan.hotel.item.sku.delete
//
// 商户对发布的宝贝套餐价格库存信息进行删除
type AlitripTuanHotelItemSkuDeleteAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelItemSkuDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTuanHotelItemSkuDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTuanHotelItemSkuDeleteAPIResponseModel).Reset()
}

// AlitripTuanHotelItemSkuDeleteAPIResponseModel is 酒店团购套餐商品SKU删除 成功返回结果
type AlitripTuanHotelItemSkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_item_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// sku列表
	TopItemSkuBaseInfoList []TopItemSkuBaseInfoList `json:"top_item_sku_base_info_list,omitempty" xml:"top_item_sku_base_info_list>top_item_sku_base_info_list,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTuanHotelItemSkuDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopItemSkuBaseInfoList = m.TopItemSkuBaseInfoList[:0]
	m.MsgCode = ""
	m.Message = ""
	m.ItemId = 0
	m.SellerId = 0
	m.Status = false
}

var poolAlitripTuanHotelItemSkuDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTuanHotelItemSkuDeleteAPIResponse)
	},
}

// GetAlitripTuanHotelItemSkuDeleteAPIResponse 从 sync.Pool 获取 AlitripTuanHotelItemSkuDeleteAPIResponse
func GetAlitripTuanHotelItemSkuDeleteAPIResponse() *AlitripTuanHotelItemSkuDeleteAPIResponse {
	return poolAlitripTuanHotelItemSkuDeleteAPIResponse.Get().(*AlitripTuanHotelItemSkuDeleteAPIResponse)
}

// ReleaseAlitripTuanHotelItemSkuDeleteAPIResponse 将 AlitripTuanHotelItemSkuDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripTuanHotelItemSkuDeleteAPIResponse(v *AlitripTuanHotelItemSkuDeleteAPIResponse) {
	v.Reset()
	poolAlitripTuanHotelItemSkuDeleteAPIResponse.Put(v)
}
