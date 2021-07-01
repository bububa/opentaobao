package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniitemItemDelete
全渠道商品删除
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作 */
func TaobaoOmniitemItemDelete(clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemDeleteAPIRequest, session string) (*omniorder.TaobaoOmniitemItemDeleteAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemItemDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
