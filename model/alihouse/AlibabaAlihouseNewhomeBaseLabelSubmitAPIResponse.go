package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomebaselabelsubmitAPIResponse 提交标签库 API返回值
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
type AlibabaalihousenewhomebaselabelsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomebaselabelsubmitAPIResponseModel
}

// AlibabaalihousenewhomebaselabelsubmitAPIResponseModel is 提交标签库 成功返回结果
type AlibabaalihousenewhomebaselabelsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_base_label_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomebaselabelsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
