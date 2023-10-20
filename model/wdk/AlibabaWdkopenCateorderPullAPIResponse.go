package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkopencateorderpullAPIResponse 商户回传餐饮加工单状态 API返回值
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
type AlibabawdkopencateorderpullAPIResponse struct {
	model.CommonResponse
	AlibabawdkopencateorderpullAPIResponseModel
}

// AlibabawdkopencateorderpullAPIResponseModel is 商户回传餐饮加工单状态 成功返回结果
type AlibabawdkopencateorderpullAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkopen_cateorder_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回
	TopBaseResult *TopBaseResult `json:"top_base_result,omitempty" xml:"top_base_result,omitempty"`
}
