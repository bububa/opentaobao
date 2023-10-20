package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscorderorderuploadAPIResponse 订单回流 API返回值
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
type AlibabaalscorderorderuploadAPIResponse struct {
	model.CommonResponse
	AlibabaalscorderorderuploadAPIResponseModel
}

// AlibabaalscorderorderuploadAPIResponseModel is 订单回流 成功返回结果
type AlibabaalscorderorderuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_order_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
