package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponse 天猫服务平台服务商查询商家投诉单 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
type TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponseModel is 天猫服务平台服务商查询商家投诉单 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_querybyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationquerybyidResult `json:"result,omitempty" xml:"result,omitempty"`
}
