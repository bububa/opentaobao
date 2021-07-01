package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterSettlementStoretransferAuditAPIRequest
新康众审批门店分账 API请求
tmall.servicecenter.settlement.storetransfer.audit

新康众审批门店分账 */
type TmallServicecenterSettlementStoretransferAuditAPIRequest struct {
	model.Params
	// 审批通过
	_auditPass bool
	// 工单id
	_workcardId int64
}

// New
