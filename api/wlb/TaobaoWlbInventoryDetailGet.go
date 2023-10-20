package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbinventorydetailget 查询库存详情
// taobao.wlb.inventory.detail.get
//
// 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
func Taobaowlbinventorydetailget(clt *core.SDKClient, req *wlb.TaobaowlbinventorydetailgetAPIRequest, session string) (*wlb.TaobaowlbinventorydetailgetAPIResponse, error) {
	var resp wlb.TaobaowlbinventorydetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
