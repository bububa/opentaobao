package servicecenter

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallServicecenterTpFundsRecoverQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTpFundsRecoverQueryAPIResponseModel).Reset()
}

// TmallServicecenterTpFundsRecoverQueryAPIResponseModel is 服务商资金权益逆向扣回的查询接口 成功返回结果
type TmallServicecenterTpFundsRecoverQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tp_funds_recover_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTpFundsRecoverQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterTpFundsRecoverQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTpFundsRecoverQueryAPIResponse)
	},
}

// GetTmallServicecenterTpFundsRecoverQueryAPIResponse 从 sync.Pool 获取 TmallServicecenterTpFundsRecoverQueryAPIResponse
func GetTmallServicecenterTpFundsRecoverQueryAPIResponse() *TmallServicecenterTpFundsRecoverQueryAPIResponse {
	return poolTmallServicecenterTpFundsRecoverQueryAPIResponse.Get().(*TmallServicecenterTpFundsRecoverQueryAPIResponse)
}

// ReleaseTmallServicecenterTpFundsRecoverQueryAPIResponse 将 TmallServicecenterTpFundsRecoverQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTpFundsRecoverQueryAPIResponse(v *TmallServicecenterTpFundsRecoverQueryAPIResponse) {
	v.Reset()
	poolTmallServicecenterTpFundsRecoverQueryAPIResponse.Put(v)
}
