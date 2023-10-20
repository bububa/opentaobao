package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillIProduct 商家查询物流商产品类型接口
// taobao.wlb.waybill.i.product
//
// 商家可以查询物流商的产品类型和服务能力。
func TaobaoWlbWaybillIProduct(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIProductAPIRequest, resp *waybill.TaobaoWlbWaybillIProductAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
