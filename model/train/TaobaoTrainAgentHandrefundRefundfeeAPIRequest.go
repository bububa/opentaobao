package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHandrefundRefundfeeAPIRequest
代理商手动退款接口 API请求
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口 */
type TaobaoTrainAgentHandrefundRefundfeeAPIRequest struct {
	model.Params
	// 主订单id
	_mainBizOrderId int64
	// 外部订单号
	_outTradeNo string
	// 退款金额,单位为分
	_refundFee int64
}

// New
