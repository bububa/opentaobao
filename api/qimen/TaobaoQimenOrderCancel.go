package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderCancel 单据取消接口
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
func TaobaoQimenOrderCancel(clt *core.SDKClient, req *qimen.TaobaoQimenOrderCancelAPIRequest, session string) (*qimen.TaobaoQimenOrderCancelAPIResponse, error) {
	var resp qimen.TaobaoQimenOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
