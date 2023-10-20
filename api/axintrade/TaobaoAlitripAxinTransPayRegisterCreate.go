package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintranspayregistercreate 提交支付服务开通
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
func Taobaoalitripaxintranspayregistercreate(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintranspayregistercreateAPIRequest, session string) (*axintrade.TaobaoalitripaxintranspayregistercreateAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintranspayregistercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
