package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlborderpageget 分页查询物流宝订单
// taobao.wlb.order.page.get
//
// 分页查询物流宝订单
func Taobaowlborderpageget(clt *core.SDKClient, req *wlb.TaobaowlborderpagegetAPIRequest, session string) (*wlb.TaobaowlborderpagegetAPIResponse, error) {
	var resp wlb.TaobaowlborderpagegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
