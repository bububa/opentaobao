package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStorecollectConsume 全渠道门店自提核销订单
// taobao.omniorder.storecollect.consume
//
// 全渠道门店自提核销订单
func TaobaoOmniorderStorecollectConsume(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStorecollectConsumeAPIRequest, resp *omniorder.TaobaoOmniorderStorecollectConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
