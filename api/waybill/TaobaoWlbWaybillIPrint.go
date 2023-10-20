package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Taobaowlbwaybilliprint 打印确认接口v1.0
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
func Taobaowlbwaybilliprint(clt *core.SDKClient, req *waybill.TaobaowlbwaybilliprintAPIRequest, session string) (*waybill.TaobaowlbwaybilliprintAPIResponse, error) {
	var resp waybill.TaobaowlbwaybilliprintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
