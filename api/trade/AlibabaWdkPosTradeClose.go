package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradeClose 轻pos品牌营销关单接口
// alibaba.wdk.pos.trade.close
//
// 轻pos品牌营销场景，提供关单接口给外部商家
func AlibabaWdkPosTradeClose(clt *core.SDKClient, req *trade.AlibabaWdkPosTradeCloseAPIRequest, resp *trade.AlibabaWdkPosTradeCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
