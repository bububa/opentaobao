package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniitemItemPublish
全渠道门店商品轻发布
taobao.omniitem.item.publish

全渠道门店商品轻发布 */
func TaobaoOmniitemItemPublish(clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemPublishAPIRequest, session string) (*omniorder.TaobaoOmniitemItemPublishAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemItemPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
