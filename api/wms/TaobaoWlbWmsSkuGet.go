package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsskuget 商品信息查询
// taobao.wlb.wms.sku.get
//
// 商品信息查询
func Taobaowlbwmsskuget(clt *core.SDKClient, req *wms.TaobaowlbwmsskugetAPIRequest, session string) (*wms.TaobaowlbwmsskugetAPIResponse, error) {
	var resp wms.TaobaowlbwmsskugetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
