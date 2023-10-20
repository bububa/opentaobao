package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbitemquery 分页查询商品
// taobao.wlb.item.query
//
// 根据状态、卖家、SKU等信息查询商品列表
func Taobaowlbitemquery(clt *core.SDKClient, req *wlb.TaobaowlbitemqueryAPIRequest, session string) (*wlb.TaobaowlbitemqueryAPIResponse, error) {
	var resp wlb.TaobaowlbitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
