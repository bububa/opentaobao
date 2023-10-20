package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcarddesensitizationqueryAPIResponse 工单查询接口 API返回值
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
type TmallservicecenterworkcarddesensitizationqueryAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcarddesensitizationqueryAPIResponseModel
}

// TmallservicecenterworkcarddesensitizationqueryAPIResponseModel is 工单查询接口 成功返回结果
type TmallservicecenterworkcarddesensitizationqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_desensitization_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *TmallservicecenterworkcarddesensitizationqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
