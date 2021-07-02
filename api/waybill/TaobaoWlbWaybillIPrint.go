package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// TaobaoWlbWaybillIPrint 打印确认接口v1.0
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
func TaobaoWlbWaybillIPrint(clt *core.SDKClient, req *waybill.TaobaoWlbWaybillIPrintAPIRequest, session string) (*waybill.TaobaoWlbWaybillIPrintAPIResponse, error) {
	var resp waybill.TaobaoWlbWaybillIPrintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
