package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbitemmapget 根据物流宝商品ID查询商品映射关系
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
func Taobaowlbitemmapget(clt *core.SDKClient, req *wlb.TaobaowlbitemmapgetAPIRequest, session string) (*wlb.TaobaowlbitemmapgetAPIResponse, error) {
	var resp wlb.TaobaowlbitemmapgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
