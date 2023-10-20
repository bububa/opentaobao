package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardQueryAPIResponse 工单查询接口 API返回值
// tmall.servicecenter.workcard.query
//
// 工单查询接口
type TmallServicecenterWorkcardQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardQueryAPIResponseModel
}

// TmallServicecenterWorkcardQueryAPIResponseModel is 工单查询接口 成功返回结果
type TmallServicecenterWorkcardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *TmallServicecenterWorkcardQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
