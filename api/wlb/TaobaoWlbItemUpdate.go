package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbitemupdate 物流宝商品修改
// taobao.wlb.item.update
//
// 修改物流宝商品信息
func Taobaowlbitemupdate(clt *core.SDKClient, req *wlb.TaobaowlbitemupdateAPIRequest, session string) (*wlb.TaobaowlbitemupdateAPIResponse, error) {
	var resp wlb.TaobaowlbitemupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
