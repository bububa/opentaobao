package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseCloseAPIResponse 服务投诉问题单关单 API返回值
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
type TmallServicecenterAnomalyrecourseCloseAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseCloseAPIResponseModel
}

// TmallServicecenterAnomalyrecourseCloseAPIResponseModel is 服务投诉问题单关单 成功返回结果
type TmallServicecenterAnomalyrecourseCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
