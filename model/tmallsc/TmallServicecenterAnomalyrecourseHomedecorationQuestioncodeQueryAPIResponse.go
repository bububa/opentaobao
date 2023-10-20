package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponse 天猫服务平台商家投诉单问题列表查询 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
type TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponseModel is 天猫服务平台商家投诉单问题列表查询 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_questioncode_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
