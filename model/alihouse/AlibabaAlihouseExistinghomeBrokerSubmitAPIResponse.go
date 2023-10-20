package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokersubmitAPIResponse 提交经纪人信息 API返回值
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
type AlibabaalihouseexistinghomebrokersubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomebrokersubmitAPIResponseModel
}

// AlibabaalihouseexistinghomebrokersubmitAPIResponseModel is 提交经纪人信息 成功返回结果
type AlibabaalihouseexistinghomebrokersubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihouseexistinghomebrokersubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
