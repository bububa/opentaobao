package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsstockinbillget 获取入库单信息
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
func Taobaowlbwmsstockinbillget(clt *core.SDKClient, req *wms.TaobaowlbwmsstockinbillgetAPIRequest, session string) (*wms.TaobaowlbwmsstockinbillgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsstockinbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
