package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemSkuUpdateAPIResponse 酒店团购套餐商品SKU更新和新增 API返回值
// alitrip.tuan.hotel.item.sku.update
//
// 商户对发布的宝贝套餐价格库存信息进行更新，对于已存在的sku，未进行传递则不会进行覆盖。skuId必须为已存在的skuId，暂不支持库存类型的更改。因发布页改造升级，2020.03.05将下线此接口的新增SKU功能，更新SKU功能将保留，但商户2020.03.05后须前往发布页进行宝贝更新后，方可调用本接口。对于日历库存宝贝日历维度的价格和库存数据的更新，此接口存在调用超时的问题，不推荐使用，若有诉求，请使用alitrip.tuan.hotel.item.sku.calendar.update接口（该接口提供增量更新能力），接口地址为https://open.taobao.com/api.htm?docId=48160&docType=2&scopeId=12326
type AlitripTuanHotelItemSkuUpdateAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelItemSkuUpdateAPIResponseModel
}

// AlitripTuanHotelItemSkuUpdateAPIResponseModel is 酒店团购套餐商品SKU更新和新增 成功返回结果
type AlitripTuanHotelItemSkuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_item_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品的sku列表
	TopItemSkuBaseInfoList []TopItemSkuBaseInfo `json:"top_item_sku_base_info_list,omitempty" xml:"top_item_sku_base_info_list>top_item_sku_base_info,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
