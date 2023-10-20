package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderCreate 退货入库单创建接口
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
func TaobaoQimenReturnorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderCreateAPIRequest, session string) (*qimen.TaobaoQimenReturnorderCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenReturnorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
