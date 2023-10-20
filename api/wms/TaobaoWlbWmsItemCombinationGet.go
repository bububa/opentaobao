package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsitemcombinationget 查询组合商品的组合关系
// taobao.wlb.wms.item.combination.get
//
// 查询组合商品的组合关系
func Taobaowlbwmsitemcombinationget(clt *core.SDKClient, req *wms.TaobaowlbwmsitemcombinationgetAPIRequest, session string) (*wms.TaobaowlbwmsitemcombinationgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsitemcombinationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
