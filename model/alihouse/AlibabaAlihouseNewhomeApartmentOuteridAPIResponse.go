package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeapartmentouteridAPIResponse 公寓更新outerid API返回值
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
type AlibabaalihousenewhomeapartmentouteridAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeapartmentouteridAPIResponseModel
}

// AlibabaalihousenewhomeapartmentouteridAPIResponseModel is 公寓更新outerid 成功返回结果
type AlibabaalihousenewhomeapartmentouteridAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_apartment_outerid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeapartmentouteridResult `json:"result,omitempty" xml:"result,omitempty"`
}
