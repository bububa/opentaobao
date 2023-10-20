package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderWorkCallbackAPIResponse 仓配作业结果回传接口 API返回值
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
type AlibabaAelophyOrderWorkCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderWorkCallbackAPIResponseModel
}

// AlibabaAelophyOrderWorkCallbackAPIResponseModel is 仓配作业结果回传接口 成功返回结果
type AlibabaAelophyOrderWorkCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_work_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传响应
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
