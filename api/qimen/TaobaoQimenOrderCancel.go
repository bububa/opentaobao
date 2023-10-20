package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenordercancel 单据取消接口
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
func Taobaoqimenordercancel(clt *core.SDKClient, req *qimen.TaobaoqimenordercancelAPIRequest, session string) (*qimen.TaobaoqimenordercancelAPIResponse, error) {
	var resp qimen.TaobaoqimenordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
