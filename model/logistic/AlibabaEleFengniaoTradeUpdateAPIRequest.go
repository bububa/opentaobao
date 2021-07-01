package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoTradeUpdateAPIRequest
更新蜂鸟扣费状态 API请求
alibaba.ele.fengniao.trade.update

汇金扣费成功后，回调该接口更新扣费状态 */
type AlibabaEleFengniaoTradeUpdateAPIRequest struct {
	model.Params
	// param 参数
	_param *Param
}

// New
