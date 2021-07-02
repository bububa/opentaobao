package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallItemsQuery 批量获取商品列表
// taobao.openmall.items.query
//
// 批量获取对联盟开放的商品列表。
func TaobaoOpenmallItemsQuery(clt *core.SDKClient, req *openmall.TaobaoOpenmallItemsQueryAPIRequest, session string) (*openmall.TaobaoOpenmallItemsQueryAPIResponse, error) {
	var resp openmall.TaobaoOpenmallItemsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
