package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabajymfulfillmentcardcallback 外部商家卡密结果回调
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
func Alibabajymfulfillmentcardcallback(clt *core.SDKClient, req *trade.AlibabajymfulfillmentcardcallbackAPIRequest, session string) (*trade.AlibabajymfulfillmentcardcallbackAPIResponse, error) {
	var resp trade.AlibabajymfulfillmentcardcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
