package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOnceTokenGet 网关一次性token获取
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
func TaobaoTopOnceTokenGet(clt *core.SDKClient, req *tbtrade.TaobaoTopOnceTokenGetAPIRequest, resp *tbtrade.TaobaoTopOnceTokenGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
