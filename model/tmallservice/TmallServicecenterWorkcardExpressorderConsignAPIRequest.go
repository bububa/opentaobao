package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardExpressorderConsignAPIRequest
物流订单呼叫运力 API请求
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力 */
type TmallServicecenterWorkcardExpressorderConsignAPIRequest struct {
	model.Params
	// 物流寄件单号（废弃）
	_expressOrderId int64
	// 工单List
	_workcardIdList []int64
	// 真实接单服务商
	_realTpNick string
	// 物流订单号
	_logisticsOrderId int64
}

// New
