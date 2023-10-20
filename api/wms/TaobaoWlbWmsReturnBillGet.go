package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsreturnbillget 获取销退收货信息
// taobao.wlb.wms.return.bill.get
//
// 通过订单号获取单个销退单收货信息。
func Taobaowlbwmsreturnbillget(clt *core.SDKClient, req *wms.TaobaowlbwmsreturnbillgetAPIRequest, session string) (*wms.TaobaowlbwmsreturnbillgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsreturnbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
