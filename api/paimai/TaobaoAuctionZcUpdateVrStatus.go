package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaoauctionzcupdatevrstatus 如视VR更新活跃状态
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
func Taobaoauctionzcupdatevrstatus(clt *core.SDKClient, req *paimai.TaobaoauctionzcupdatevrstatusAPIRequest, session string) (*paimai.TaobaoauctionzcupdatevrstatusAPIResponse, error) {
	var resp paimai.TaobaoauctionzcupdatevrstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
