package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbordercreate 创建物流宝订单
// taobao.wlb.order.create
//
// 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
func Taobaowlbordercreate(clt *core.SDKClient, req *wlb.TaobaowlbordercreateAPIRequest, session string) (*wlb.TaobaowlbordercreateAPIResponse, error) {
	var resp wlb.TaobaowlbordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
