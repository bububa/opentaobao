package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWtTradeOrderResultcallbackAPIRequest
商家回调接口 API请求
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货 */
type TaobaoWtTradeOrderResultcallbackAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *OrderResultDto
}

// New
