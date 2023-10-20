package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniOrderGoodsReady 备货完成
// taobao.omni.order.goods.ready
//
// 备货完成
func TaobaoOmniOrderGoodsReady(clt *core.SDKClient, req *omniorder.TaobaoOmniOrderGoodsReadyAPIRequest, session string) (*omniorder.TaobaoOmniOrderGoodsReadyAPIResponse, error) {
	var resp omniorder.TaobaoOmniOrderGoodsReadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
