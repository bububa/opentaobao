package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemMapGet 根据物流宝商品ID查询商品映射关系
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
func TaobaoWlbItemMapGet(clt *core.SDKClient, req *wlb.TaobaoWlbItemMapGetAPIRequest, session string) (*wlb.TaobaoWlbItemMapGetAPIResponse, error) {
	var resp wlb.TaobaoWlbItemMapGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
