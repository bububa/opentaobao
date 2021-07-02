package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelItemInfoGetAPIResponse 宝贝信息查询接口 API返回值
// alitrip.tuan.hotel.item.info.get
//
// 商家查询发布的宝贝详情信息
type AlitripTuanHotelItemInfoGetAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelItemInfoGetAPIResponseModel
}

// AlitripTuanHotelItemInfoGetAPIResponseModel is 宝贝信息查询接口 成功返回结果
type AlitripTuanHotelItemInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_item_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 宝贝基本信息
	ItemInfo *TuanItemSellParamVO `json:"item_info,omitempty" xml:"item_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// sku列表
	TuanItemSkuList []TopTuanItemSkuVOList `json:"tuan_item_sku_list,omitempty" xml:"tuan_item_sku_list>top_tuan_item_sku_vo_list,omitempty"`
	// 电子核销库信息，日历库存和国际套餐暂不支持电子凭证，则无数值返回
	TuanEticketPackage *TuanEticketPackageVo `json:"tuan_eticket_package,omitempty" xml:"tuan_eticket_package,omitempty"`
	// 关联门店列表
	Stores []TopStoreVO `json:"stores,omitempty" xml:"stores>top_store_vo,omitempty"`
	// 关联POI列表
	RelatedPoiDetailVOList []RelatedPoiDetailVo `json:"related_poi_detail_v_o_list,omitempty" xml:"related_poi_detail_v_o_list>related_poi_detail_vo,omitempty"`
	// 关联礼包列表
	TuanItemRelateGiftList []TuanItemRelateGiftVO `json:"tuan_item_relate_gift_list,omitempty" xml:"tuan_item_relate_gift_list>tuan_item_relate_gift_vo,omitempty"`
}
