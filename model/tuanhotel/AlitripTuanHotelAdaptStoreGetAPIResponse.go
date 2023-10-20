package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelAdaptStoreGetAPIResponse 酒店团购套餐关联适用门店 API返回值
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
type AlitripTuanHotelAdaptStoreGetAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelAdaptStoreGetAPIResponseModel
}

// AlitripTuanHotelAdaptStoreGetAPIResponseModel is 酒店团购套餐关联适用门店 成功返回结果
type AlitripTuanHotelAdaptStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_adapt_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关联门店列表
	StoreDetailList []StoreDetailVoList `json:"store_detail_list,omitempty" xml:"store_detail_list>store_detail_vo_list,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
