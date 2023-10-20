package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsconsignbillget 获取销售订单发货信息
// taobao.wlb.wms.consign.bill.get
//
// 获取销售订单发货信息
func Taobaowlbwmsconsignbillget(clt *core.SDKClient, req *wms.TaobaowlbwmsconsignbillgetAPIRequest, session string) (*wms.TaobaowlbwmsconsignbillgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsconsignbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
