package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoservindustryfinancegeexorderupdate 即科订单结果更新回调
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
func Taobaoservindustryfinancegeexorderupdate(clt *core.SDKClient, req *trade.TaobaoservindustryfinancegeexorderupdateAPIRequest, session string) (*trade.TaobaoservindustryfinancegeexorderupdateAPIResponse, error) {
	var resp trade.TaobaoservindustryfinancegeexorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
