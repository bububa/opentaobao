package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillIQuerydetail 查面单号状态v1.0
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
func TaobaoWlbWaybillIQuerydetail(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIQuerydetailAPIRequest, resp *waybill.TaobaoWlbWaybillIQuerydetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
