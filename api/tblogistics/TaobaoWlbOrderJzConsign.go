package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaowlborderjzconsign 家装发货接口
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
func Taobaowlborderjzconsign(clt *core.SDKClient, req *tblogistics.TaobaowlborderjzconsignAPIRequest, session string) (*tblogistics.TaobaowlborderjzconsignAPIResponse, error) {
	var resp tblogistics.TaobaowlborderjzconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
