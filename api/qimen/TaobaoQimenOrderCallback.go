package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderCallback 配送拦截接口
// taobao.qimen.order.callback
//
// 配送拦截
func TaobaoQimenOrderCallback(clt *core.SDKClient, req *qimen.TaobaoQimenOrderCallbackAPIRequest, resp *qimen.TaobaoQimenOrderCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
