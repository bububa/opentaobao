package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationappealAPIResponse 天猫服务平台商家投诉单服务商申诉接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
type TmallservicecenteranomalyrecoursehomedecorationappealAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationappealAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationappealAPIResponseModel is 天猫服务平台商家投诉单服务商申诉接口 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationappealAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_appeal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationappealResult `json:"result,omitempty" xml:"result,omitempty"`
}
