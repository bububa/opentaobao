package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest
退票通知 API请求
taobao.train.agent.returnticket.confirm.vtwo

火车票代理商接口——退票通知回调 */
type TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest struct {
	model.Params
	// 用户id
	_buyerId int64
	// 是否同意退票
	_agreeReturn bool
	// 退票金额，单位分
	_refundFee int64
	// 淘宝主订单id
	_mainBizOrderId int64
	// 代理商id
	_agentId int64
	// 拒绝退票原因
	_refuseReturnReason string
	// 淘宝子订单id
	_subBizOrderId int64
	// 关闭通知用户 true:关闭 false:不关闭
	_closeRefundNotify bool
}

// New
