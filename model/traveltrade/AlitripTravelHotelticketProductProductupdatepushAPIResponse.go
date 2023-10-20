package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketproductproductupdatepushAPIResponse 产品批量变更推送通知 API返回值
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
type AlitriptravelhotelticketproductproductupdatepushAPIResponse struct {
	model.CommonResponse
	AlitriptravelhotelticketproductproductupdatepushAPIResponseModel
}

// AlitriptravelhotelticketproductproductupdatepushAPIResponseModel is 产品批量变更推送通知 成功返回结果
type AlitriptravelhotelticketproductproductupdatepushAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_hotelticket_product_productupdatepush_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
