package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOnlineCancelAPIRequest
取消物流订单接口 API请求
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。 */
type TaobaoLogisticsOnlineCancelAPIRequest struct {
	model.Params
	// 淘宝交易ID
	_tid int64
}

// New
