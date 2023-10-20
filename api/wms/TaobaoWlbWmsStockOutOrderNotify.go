package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsstockoutordernotify 出库单通知
// taobao.wlb.wms.stock.out.order.notify
//
// 出库单通知
func Taobaowlbwmsstockoutordernotify(clt *core.SDKClient, req *wms.TaobaowlbwmsstockoutordernotifyAPIRequest, session string) (*wms.TaobaowlbwmsstockoutordernotifyAPIResponse, error) {
	var resp wms.TaobaowlbwmsstockoutordernotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
