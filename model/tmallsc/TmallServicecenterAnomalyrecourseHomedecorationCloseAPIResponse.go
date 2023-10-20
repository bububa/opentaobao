package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponse 天猫服务平台商家投诉单服务商完结接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
type TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponseModel is 天猫服务平台商家投诉单服务商完结接口 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationcloseResult `json:"result,omitempty" xml:"result,omitempty"`
}
