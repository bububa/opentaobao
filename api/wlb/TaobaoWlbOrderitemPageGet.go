package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlborderitempageget 分页查询物流宝订单商品详情
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
func Taobaowlborderitempageget(clt *core.SDKClient, req *wlb.TaobaowlborderitempagegetAPIRequest, session string) (*wlb.TaobaowlborderitempagegetAPIResponse, error) {
	var resp wlb.TaobaowlborderitempagegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
