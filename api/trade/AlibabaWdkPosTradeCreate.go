package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradeCreate 轻pos品牌营销下单接口
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
func AlibabaWdkPosTradeCreate(clt *core.SDKClient, req *trade.AlibabaWdkPosTradeCreateAPIRequest, resp *trade.AlibabaWdkPosTradeCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
