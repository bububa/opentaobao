package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponse 天猫服务平台商家投诉单服务商响应接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
type TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponseModel is 天猫服务平台商家投诉单服务商响应接口 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_response_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationresponseResult `json:"result,omitempty" xml:"result,omitempty"`
}
