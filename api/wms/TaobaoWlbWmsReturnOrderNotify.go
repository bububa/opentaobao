package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsreturnordernotify 销售退货通知
// taobao.wlb.wms.return.order.notify
//
// 销售退货通知
func Taobaowlbwmsreturnordernotify(clt *core.SDKClient, req *wms.TaobaowlbwmsreturnordernotifyAPIRequest, session string) (*wms.TaobaowlbwmsreturnordernotifyAPIResponse, error) {
	var resp wms.TaobaowlbwmsreturnordernotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
