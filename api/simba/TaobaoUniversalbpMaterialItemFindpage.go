package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpmaterialitemfindpage 分页查询商品信息
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
func Taobaouniversalbpmaterialitemfindpage(clt *core.SDKClient, req *simba.TaobaouniversalbpmaterialitemfindpageAPIRequest, session string) (*simba.TaobaouniversalbpmaterialitemfindpageAPIResponse, error) {
	var resp simba.TaobaouniversalbpmaterialitemfindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
