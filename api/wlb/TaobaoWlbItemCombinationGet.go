package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbitemcombinationget 根据商品id查询商品组合关系
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
func Taobaowlbitemcombinationget(clt *core.SDKClient, req *wlb.TaobaowlbitemcombinationgetAPIRequest, session string) (*wlb.TaobaowlbitemcombinationgetAPIResponse, error) {
	var resp wlb.TaobaowlbitemcombinationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
