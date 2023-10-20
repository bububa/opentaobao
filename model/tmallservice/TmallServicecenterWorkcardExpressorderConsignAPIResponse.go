package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExpressorderConsignAPIResponse 物流订单呼叫运力 API返回值
// tmall.servicecenter.workcard.expressorder.consign
//
// 天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
type TmallServicecenterWorkcardExpressorderConsignAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardExpressorderConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExpressorderConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardExpressorderConsignAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardExpressorderConsignAPIResponseModel is 物流订单呼叫运力 成功返回结果
type TmallServicecenterWorkcardExpressorderConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_expressorder_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardExpressorderConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardExpressorderConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardExpressorderConsignAPIResponse)
	},
}

// GetTmallServicecenterWorkcardExpressorderConsignAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardExpressorderConsignAPIResponse
func GetTmallServicecenterWorkcardExpressorderConsignAPIResponse() *TmallServicecenterWorkcardExpressorderConsignAPIResponse {
	return poolTmallServicecenterWorkcardExpressorderConsignAPIResponse.Get().(*TmallServicecenterWorkcardExpressorderConsignAPIResponse)
}

// ReleaseTmallServicecenterWorkcardExpressorderConsignAPIResponse 将 TmallServicecenterWorkcardExpressorderConsignAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardExpressorderConsignAPIResponse(v *TmallServicecenterWorkcardExpressorderConsignAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardExpressorderConsignAPIResponse.Put(v)
}
