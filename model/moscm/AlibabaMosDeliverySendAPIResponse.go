package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosdeliverysendAPIResponse 发货 API返回值
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
type AlibabamosdeliverysendAPIResponse struct {
	model.CommonResponse
	AlibabamosdeliverysendAPIResponseModel
}

// AlibabamosdeliverysendAPIResponseModel is 发货 成功返回结果
type AlibabamosdeliverysendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_delivery_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamosdeliverysendResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
