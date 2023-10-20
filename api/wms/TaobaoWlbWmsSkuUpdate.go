package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsskuupdate 商品信息的更新
// taobao.wlb.wms.sku.update
//
// 商品信息的更新
func Taobaowlbwmsskuupdate(clt *core.SDKClient, req *wms.TaobaowlbwmsskuupdateAPIRequest, session string) (*wms.TaobaowlbwmsskuupdateAPIResponse, error) {
	var resp wms.TaobaowlbwmsskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
