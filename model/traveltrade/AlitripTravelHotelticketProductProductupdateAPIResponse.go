package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketProductProductupdateAPIResponse 产品批量变更通知 API返回值
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
type AlitripTravelHotelticketProductProductupdateAPIResponse struct {
	model.CommonResponse
	AlitripTravelHotelticketProductProductupdateAPIResponseModel
}

// AlitripTravelHotelticketProductProductupdateAPIResponseModel is 产品批量变更通知 成功返回结果
type AlitripTravelHotelticketProductProductupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_product_productupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
