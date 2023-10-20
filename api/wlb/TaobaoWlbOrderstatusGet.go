package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlborderstatusget 物流宝订单流转状态查询
// taobao.wlb.orderstatus.get
//
// 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
func Taobaowlborderstatusget(clt *core.SDKClient, req *wlb.TaobaowlborderstatusgetAPIRequest, session string) (*wlb.TaobaowlborderstatusgetAPIResponse, error) {
	var resp wlb.TaobaowlborderstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
