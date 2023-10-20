package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTpFundsSendQueryAPIResponse 服务商资金权益发放的查询接口 API返回值
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
type TmallServicecenterTpFundsSendQueryAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTpFundsSendQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterTpFundsSendQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTpFundsSendQueryAPIResponseModel).Reset()
}

// TmallServicecenterTpFundsSendQueryAPIResponseModel is 服务商资金权益发放的查询接口 成功返回结果
type TmallServicecenterTpFundsSendQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tp_funds_send_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTpFundsSendQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterTpFundsSendQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTpFundsSendQueryAPIResponse)
	},
}

// GetTmallServicecenterTpFundsSendQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterTpFundsSendQueryAPIResponse
func GetTmallServicecenterTpFundsSendQueryAPIResponse() *TmallServicecenterTpFundsSendQueryAPIResponse {
	return poolTmallServicecenterTpFundsSendQueryAPIResponse.Get().(*TmallServicecenterTpFundsSendQueryAPIResponse)
}

// ReleaseTmallServicecenterTpFundsSendQueryAPIResponse 将 TmallServicecenterTpFundsSendQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTpFundsSendQueryAPIResponse(v *TmallServicecenterTpFundsSendQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterTpFundsSendQueryAPIResponse.Put(v)
}
