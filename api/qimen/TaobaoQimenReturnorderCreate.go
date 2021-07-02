package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderCreate 退货入库单创建接口
// taobao.qimen.returnorder.create
//
// ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
func TaobaoQimenReturnorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderCreateAPIRequest, session string) (*qimen.TaobaoQimenReturnorderCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenReturnorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
