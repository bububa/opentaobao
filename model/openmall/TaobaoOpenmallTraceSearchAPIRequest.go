package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTraceSearchAPIRequest
获取Openmall订单物流流转信息 API请求
taobao.openmall.trace.search

获取Openmall订单物流流转信息 */
type TaobaoOpenmallTraceSearchAPIRequest struct {
	model.Params
	// 签约支付宝代扣的账号
	_distributor string
	// 淘宝订单编号
	_tid int64
}

// New
