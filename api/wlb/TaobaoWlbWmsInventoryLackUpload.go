package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbwmsinventorylackupload 缺货通知
// taobao.wlb.wms.inventory.lack.upload
//
// 缺货通知
func Taobaowlbwmsinventorylackupload(clt *core.SDKClient, req *wlb.TaobaowlbwmsinventorylackuploadAPIRequest, session string) (*wlb.TaobaowlbwmsinventorylackuploadAPIResponse, error) {
	var resp wlb.TaobaowlbwmsinventorylackuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
