package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniitemClassifyStoreQuery
根据门店查分类信息
taobao.omniitem.classify.store.query

根据门店查分类信息 */
func TaobaoOmniitemClassifyStoreQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyStoreQueryAPIRequest, session string) (*omniorder.TaobaoOmniitemClassifyStoreQueryAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemClassifyStoreQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
