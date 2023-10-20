package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderdtdconsign 门店自送发货
// taobao.omniorder.dtd.consign
//
// 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
func Taobaoomniorderdtdconsign(clt *core.SDKClient, req *omniorder.TaobaoomniorderdtdconsignAPIRequest, session string) (*omniorder.TaobaoomniorderdtdconsignAPIResponse, error) {
	var resp omniorder.TaobaoomniorderdtdconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
