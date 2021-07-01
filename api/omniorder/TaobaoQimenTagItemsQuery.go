package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoQimenTagItemsQuery
打标结果查询-标维度
taobao.qimen.tag.items.query

调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。 */
func TaobaoQimenTagItemsQuery(clt *core.SDKClient, req *omniorder.TaobaoQimenTagItemsQueryAPIRequest, session string) (*omniorder.TaobaoQimenTagItemsQueryAPIResponse, error) {
	var resp omniorder.TaobaoQimenTagItemsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
