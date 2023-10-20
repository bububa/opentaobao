package icbulogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddresscitylistAPIResponse 四级地址库-市 API返回值
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
type AlibabaonetouchlogisticsexpressaddresscitylistAPIResponse struct {
	model.CommonResponse
	AlibabaonetouchlogisticsexpressaddresscitylistAPIResponseModel
}

// AlibabaonetouchlogisticsexpressaddresscitylistAPIResponseModel is 四级地址库-市 成功返回结果
type AlibabaonetouchlogisticsexpressaddresscitylistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_city_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaonetouchlogisticsexpressaddresscitylistResult `json:"result,omitempty" xml:"result,omitempty"`
}
