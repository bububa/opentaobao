package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemClassifyItemQuery 根据分类查商品信息
// taobao.omniitem.classify.item.query
//
// 商家根据分类查商品
func TaobaoOmniitemClassifyItemQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniitemClassifyItemQueryAPIRequest, session string) (*omniorder.TaobaoOmniitemClassifyItemQueryAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemClassifyItemQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
