package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbTmsorderQuery 通过物流订单编号查询物流信息
// taobao.wlb.tmsorder.query
//
// 通过物流订单编号分页查询物流信息
func TaobaoWlbTmsorderQuery(clt *core.SDKClient, req *wlb.TaobaoWlbTmsorderQueryAPIRequest, resp *wlb.TaobaoWlbTmsorderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
