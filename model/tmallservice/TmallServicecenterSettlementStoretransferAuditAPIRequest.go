package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentersettlementstoretransferauditAPIRequest 新康众审批门店分账 API请求
// tmall.servicecenter.settlement.storetransfer.audit
//
// 新康众审批门店分账
type TmallservicecentersettlementstoretransferauditAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 审批通过
	_auditPass bool
}

// NewTmallservicecentersettlementstoretransferauditRequest 初始化TmallservicecentersettlementstoretransferauditAPIRequest对象
func NewTmallservicecentersettlementstoretransferauditRequest() *TmallservicecentersettlementstoretransferauditAPIRequest {
	return &TmallservicecentersettlementstoretransferauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentersettlementstoretransferauditAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.settlement.storetransfer.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentersettlementstoretransferauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentersettlementstoretransferauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecentersettlementstoretransferauditAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecentersettlementstoretransferauditAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetAuditPass is AuditPass Setter
// 审批通过
func (r *TmallservicecentersettlementstoretransferauditAPIRequest) SetAuditPass(_auditPass bool) error {
	r._auditPass = _auditPass
	r.Set("audit_pass", _auditPass)
	return nil
}

// GetAuditPass AuditPass Getter
func (r TmallservicecentersettlementstoretransferauditAPIRequest) GetAuditPass() bool {
	return r._auditPass
}
