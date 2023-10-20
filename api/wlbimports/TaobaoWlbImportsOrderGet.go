package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportsorderget 物流订单获取
// taobao.wlb.imports.order.get
//
// 一般进口物流订单获取
func Taobaowlbimportsorderget(clt *core.SDKClient, req *wlbimports.TaobaowlbimportsordergetAPIRequest, session string) (*wlbimports.TaobaowlbimportsordergetAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportsordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
