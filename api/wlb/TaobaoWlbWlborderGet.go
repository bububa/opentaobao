package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbwlborderget 根据物流宝订单编号查询物流宝订单概要信息
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
func Taobaowlbwlborderget(clt *core.SDKClient, req *wlb.TaobaowlbwlbordergetAPIRequest, session string) (*wlb.TaobaowlbwlbordergetAPIResponse, error) {
	var resp wlb.TaobaowlbwlbordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
