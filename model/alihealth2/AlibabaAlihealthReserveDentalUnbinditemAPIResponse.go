package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthreservedentalunbinditemAPIResponse 解绑商品信息 API返回值
// alibaba.alihealth.reserve.dental.unbinditem
//
// 绑定门店信息，商品信息
type AlibabaalihealthreservedentalunbinditemAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthreservedentalunbinditemAPIResponseModel
}

// AlibabaalihealthreservedentalunbinditemAPIResponseModel is 解绑商品信息 成功返回结果
type AlibabaalihealthreservedentalunbinditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_reserve_dental_unbinditem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihealthreservedentalunbinditemResult `json:"result,omitempty" xml:"result,omitempty"`
}
