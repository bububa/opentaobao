package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomercchangestatusAPIResponse 图文草稿状态更新 API返回值
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
type AlibabaalihousenewhomercchangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomercchangestatusAPIResponseModel
}

// AlibabaalihousenewhomercchangestatusAPIResponseModel is 图文草稿状态更新 成功返回结果
type AlibabaalihousenewhomercchangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomercchangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
