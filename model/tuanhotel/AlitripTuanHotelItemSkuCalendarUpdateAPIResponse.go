package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemSkuCalendarUpdateAPIResponse 酒店非标套餐商品日历库存宝贝SKU更新接口 API返回值
// alitrip.tuan.hotel.item.sku.calendar.update
//
// 商户对发布的日历库存类型的宝贝套餐价格库存信息进行更新，仅提供日历库存的宝贝SKU的更新功能，skuId须传递商品已存在的skuId，若想进行SKU新增操作，请选择使用alitrip.tuan.hotel.item.sku.update接口。提供增量更新SKU功能，对于日历库存若传递日期信息，参数中若包含某一日期的价格和库存，则对此日期的数据进行覆盖更新，若不传递则保留此日期的价格库存信息。
type AlitripTuanHotelItemSkuCalendarUpdateAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelItemSkuCalendarUpdateAPIResponseModel
}

// AlitripTuanHotelItemSkuCalendarUpdateAPIResponseModel is 酒店非标套餐商品日历库存宝贝SKU更新接口 成功返回结果
type AlitripTuanHotelItemSkuCalendarUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_item_sku_calendar_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品的sku列表
	TopItemSkuBaseInfoList []TopItemSkuBaseInfo `json:"top_item_sku_base_info_list,omitempty" xml:"top_item_sku_base_info_list>top_item_sku_base_info,omitempty"`
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
