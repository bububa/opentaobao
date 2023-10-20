package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsBind 交易开放商品绑定
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
func TaobaoOpentradeToolsItemsBind(clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsBindAPIRequest, session string) (*opentrade.TaobaoOpentradeToolsItemsBindAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeToolsItemsBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
