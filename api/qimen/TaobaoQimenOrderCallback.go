package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenordercallback 配送拦截接口
// taobao.qimen.order.callback
//
// 配送拦截
func Taobaoqimenordercallback(clt *core.SDKClient, req *qimen.TaobaoqimenordercallbackAPIRequest, session string) (*qimen.TaobaoqimenordercallbackAPIResponse, error) {
	var resp qimen.TaobaoqimenordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
