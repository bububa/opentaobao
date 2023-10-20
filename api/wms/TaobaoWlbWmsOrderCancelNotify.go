package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsordercancelnotify 单据取消接口
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
func Taobaowlbwmsordercancelnotify(clt *core.SDKClient, req *wms.TaobaowlbwmsordercancelnotifyAPIRequest, session string) (*wms.TaobaowlbwmsordercancelnotifyAPIResponse, error) {
	var resp wms.TaobaowlbwmsordercancelnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
