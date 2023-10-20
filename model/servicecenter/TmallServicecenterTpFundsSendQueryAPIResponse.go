package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertpfundssendqueryAPIResponse 服务商资金权益发放的查询接口 API返回值
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
type TmallservicecentertpfundssendqueryAPIResponse struct {
	model.CommonResponse
	TmallservicecentertpfundssendqueryAPIResponseModel
}

// TmallservicecentertpfundssendqueryAPIResponseModel is 服务商资金权益发放的查询接口 成功返回结果
type TmallservicecentertpfundssendqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tp_funds_send_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
