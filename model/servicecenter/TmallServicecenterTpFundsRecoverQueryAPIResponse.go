package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTpFundsRecoverQueryAPIResponse 服务商资金权益逆向扣回的查询接口 API返回值
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
type TmallServicecenterTpFundsRecoverQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTpFundsRecoverQueryAPIResponseModel
}

// TmallServicecenterTpFundsRecoverQueryAPIResponseModel is 服务商资金权益逆向扣回的查询接口 成功返回结果
type TmallServicecenterTpFundsRecoverQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tp_funds_recover_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
