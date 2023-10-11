package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse 天猫服务平台服务商代商家发起投诉单 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
type TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel
}

// TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel is 天猫服务平台服务商代商家发起投诉单 成功返回结果
type TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallServicecenterAnomalyrecourseHomedecorationCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
