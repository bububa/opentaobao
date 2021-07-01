package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderCallbackAPIRequest
配送拦截接口 API请求
taobao.qimen.order.callback

配送拦截 */
type TaobaoQimenOrderCallbackAPIRequest struct {
	model.Params
	//
	_request *OrderCallbackRequestDo
}

// New
