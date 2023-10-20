package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemdelete 删除单条商品
// taobao.item.delete
//
// 删除单条商品
func Taobaoitemdelete(clt *core.SDKClient, req *tbitem.TaobaoitemdeleteAPIRequest, session string) (*tbitem.TaobaoitemdeleteAPIResponse, error) {
	var resp tbitem.TaobaoitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
