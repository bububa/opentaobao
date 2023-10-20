package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportsordercancel 一般进口取消物流订单
// taobao.wlb.imports.order.cancel
//
// 商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
func Taobaowlbimportsordercancel(clt *core.SDKClient, req *wlbimports.TaobaowlbimportsordercancelAPIRequest, session string) (*wlbimports.TaobaowlbimportsordercancelAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportsordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
