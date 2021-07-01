package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* AlibabaWdkPosTradeReverse
轻pos品牌营销退款接口
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口 */
func AlibabaWdkPosTradeReverse(clt *core.SDKClient, req *trade.AlibabaWdkPosTradeReverseAPIRequest, session string) (*trade.AlibabaWdkPosTradeReverseAPIResponse, error) {
	var resp trade.AlibabaWdkPosTradeReverseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
