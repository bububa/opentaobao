package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabajymfulfillmentioschargecallback 代充充值回调
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
func Alibabajymfulfillmentioschargecallback(clt *core.SDKClient, req *trade.AlibabajymfulfillmentioschargecallbackAPIRequest, session string) (*trade.AlibabajymfulfillmentioschargecallbackAPIResponse, error) {
	var resp trade.AlibabajymfulfillmentioschargecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
