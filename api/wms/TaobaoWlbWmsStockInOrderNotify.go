package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsstockinordernotify 入库通知单
// taobao.wlb.wms.stock.in.order.notify
//
// 入库通知单
func Taobaowlbwmsstockinordernotify(clt *core.SDKClient, req *wms.TaobaowlbwmsstockinordernotifyAPIRequest, session string) (*wms.TaobaowlbwmsstockinordernotifyAPIResponse, error) {
	var resp wms.TaobaowlbwmsstockinordernotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
