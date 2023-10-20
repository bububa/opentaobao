package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniOrderGoodsReady 备货完成
// taobao.omni.order.goods.ready
//
// 备货完成
func TaobaoOmniOrderGoodsReady(clt *core.SDKClient, req *omniorder.TaobaoOmniOrderGoodsReadyAPIRequest, resp *omniorder.TaobaoOmniOrderGoodsReadyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
