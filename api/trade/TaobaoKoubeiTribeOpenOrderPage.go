package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaokoubeitribeopenorderpage 口碑综合体订单列表信息查询
// taobao.koubei.tribe.open.order.page
//
// 查询口碑商圈用户的订单列表信息
func Taobaokoubeitribeopenorderpage(clt *core.SDKClient, req *trade.TaobaokoubeitribeopenorderpageAPIRequest, session string) (*trade.TaobaokoubeitribeopenorderpageAPIResponse, error) {
	var resp trade.TaobaokoubeitribeopenorderpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
