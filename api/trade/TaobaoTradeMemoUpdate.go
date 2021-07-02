package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeMemoUpdate 修改交易备注
// taobao.trade.memo.update
//
// 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
func TaobaoTradeMemoUpdate(clt *core.SDKClient, req *trade.TaobaoTradeMemoUpdateAPIRequest, session string) (*trade.TaobaoTradeMemoUpdateAPIResponse, error) {
	var resp trade.TaobaoTradeMemoUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
