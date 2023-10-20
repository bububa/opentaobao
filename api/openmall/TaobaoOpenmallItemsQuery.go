package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallitemsquery 批量获取商品列表
// taobao.openmall.items.query
//
// 批量获取对联盟开放的商品列表。
func Taobaoopenmallitemsquery(clt *core.SDKClient, req *openmall.TaobaoopenmallitemsqueryAPIRequest, session string) (*openmall.TaobaoopenmallitemsqueryAPIResponse, error) {
	var resp openmall.TaobaoopenmallitemsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
