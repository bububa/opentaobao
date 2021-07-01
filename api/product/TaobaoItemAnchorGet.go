package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemAnchorGet
获取可用宝贝描述规范化模块
taobao.item.anchor.get

根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点 */
func TaobaoItemAnchorGet(clt *core.SDKClient, req *product.TaobaoItemAnchorGetAPIRequest, session string) (*product.TaobaoItemAnchorGetAPIResponse, error) {
	var resp product.TaobaoItemAnchorGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
