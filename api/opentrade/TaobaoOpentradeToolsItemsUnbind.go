package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsUnbind 交易开放商品解绑
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
func TaobaoOpentradeToolsItemsUnbind(clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsUnbindAPIRequest, session string) (*opentrade.TaobaoOpentradeToolsItemsUnbindAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeToolsItemsUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
