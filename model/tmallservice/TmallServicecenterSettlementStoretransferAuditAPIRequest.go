package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSettlementStoretransferAuditAPIRequest 新康众审批门店分账 API请求
// tmall.servicecenter.settlement.storetransfer.audit
//
// 新康众审批门店分账
type TmallServicecenterSettlementStoretransferAuditAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 审批通过
	_auditPass bool
}

// NewTmallServicecenterSettlementStoretransferAuditRequest 初始化TmallServicecenterSettlementStoretransferAuditAPIRequest对象
func NewTmallServicecenterSettlementStoretransferAuditRequest() *TmallServicecenterSettlementStoretransferAuditAPIRequest {
	return &TmallServicecenterSettlementStoretransferAuditAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterSettlementStoretransferAuditAPIRequest) Reset() {
	r._workcardId = 0
	r._auditPass = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.settlement.storetransfer.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterSettlementStoretransferAuditAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetAuditPass is AuditPass Setter
// 审批通过
func (r *TmallServicecenterSettlementStoretransferAuditAPIRequest) SetAuditPass(_auditPass bool) error {
	r._auditPass = _auditPass
	r.Set("audit_pass", _auditPass)
	return nil
}

// GetAuditPass AuditPass Getter
func (r TmallServicecenterSettlementStoretransferAuditAPIRequest) GetAuditPass() bool {
	return r._auditPass
}

var poolTmallServicecenterSettlementStoretransferAuditAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterSettlementStoretransferAuditRequest()
	},
}

// GetTmallServicecenterSettlementStoretransferAuditRequest 从 sync.Pool 获取 TmallServicecenterSettlementStoretransferAuditAPIRequest
func GetTmallServicecenterSettlementStoretransferAuditAPIRequest() *TmallServicecenterSettlementStoretransferAuditAPIRequest {
	return poolTmallServicecenterSettlementStoretransferAuditAPIRequest.Get().(*TmallServicecenterSettlementStoretransferAuditAPIRequest)
}

// ReleaseTmallServicecenterSettlementStoretransferAuditAPIRequest 将 TmallServicecenterSettlementStoretransferAuditAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterSettlementStoretransferAuditAPIRequest(v *TmallServicecenterSettlementStoretransferAuditAPIRequest) {
	v.Reset()
	poolTmallServicecenterSettlementStoretransferAuditAPIRequest.Put(v)
}
