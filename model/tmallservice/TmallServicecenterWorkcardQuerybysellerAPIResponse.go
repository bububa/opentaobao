package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardquerybysellerAPIResponse 工单查询接口（面向商家） API返回值
// tmall.servicecenter.workcard.querybyseller
//
// 查询工单
type TmallservicecenterworkcardquerybysellerAPIResponse struct {
	model.CommonResponse
	TmallservicecenterworkcardquerybysellerAPIResponseModel
}

// TmallservicecenterworkcardquerybysellerAPIResponseModel is 工单查询接口（面向商家） 成功返回结果
type TmallservicecenterworkcardquerybysellerAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_querybyseller_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallservicecenterworkcardquerybysellerResult `json:"result,omitempty" xml:"result,omitempty"`
}
