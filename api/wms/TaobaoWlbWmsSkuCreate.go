package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsskucreate 商品同步
// taobao.wlb.wms.sku.create
//
// 商品同步
func Taobaowlbwmsskucreate(clt *core.SDKClient, req *wms.TaobaowlbwmsskucreateAPIRequest, session string) (*wms.TaobaowlbwmsskucreateAPIResponse, error) {
	var resp wms.TaobaowlbwmsskucreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
