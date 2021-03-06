package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardCancelAPIResponse 服务平台工单取消接口 API返回值
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
type AlibabaServicecenterWorkcardCancelAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardCancelAPIResponseModel
}

// AlibabaServicecenterWorkcardCancelAPIResponseModel is 服务平台工单取消接口 成功返回结果
type AlibabaServicecenterWorkcardCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
