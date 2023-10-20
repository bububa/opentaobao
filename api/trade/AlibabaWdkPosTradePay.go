package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdkpostradepay 轻pos品牌营销支付接口
// alibaba.wdk.pos.trade.pay
//
// 轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
func Alibabawdkpostradepay(clt *core.SDKClient, req *trade.AlibabawdkpostradepayAPIRequest, session string) (*trade.AlibabawdkpostradepayAPIResponse, error) {
	var resp trade.AlibabawdkpostradepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
