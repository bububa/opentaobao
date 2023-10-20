package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSettlementStoretransferAuditAPIResponse 新康众审批门店分账 API返回值
// tmall.servicecenter.settlement.storetransfer.audit
//
// 新康众审批门店分账
type TmallServicecenterSettlementStoretransferAuditAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSettlementStoretransferAuditAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterSettlementStoretransferAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterSettlementStoretransferAuditAPIResponseModel).Reset()
}

// TmallServicecenterSettlementStoretransferAuditAPIResponseModel is 新康众审批门店分账 成功返回结果
type TmallServicecenterSettlementStoretransferAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_settlement_storetransfer_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分账审批通知结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterSettlementStoretransferAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterSettlementStoretransferAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterSettlementStoretransferAuditAPIResponse)
	},
}

// GetTmallServicecenterSettlementStoretransferAuditAPIResponse 从 sync.Pool 获取 TmallServicecenterSettlementStoretransferAuditAPIResponse
func GetTmallServicecenterSettlementStoretransferAuditAPIResponse() *TmallServicecenterSettlementStoretransferAuditAPIResponse {
	return poolTmallServicecenterSettlementStoretransferAuditAPIResponse.Get().(*TmallServicecenterSettlementStoretransferAuditAPIResponse)
}

// ReleaseTmallServicecenterSettlementStoretransferAuditAPIResponse 将 TmallServicecenterSettlementStoretransferAuditAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterSettlementStoretransferAuditAPIResponse(v *TmallServicecenterSettlementStoretransferAuditAPIResponse) {
	v.Reset()
	poolTmallServicecenterSettlementStoretransferAuditAPIResponse.Put(v)
}
