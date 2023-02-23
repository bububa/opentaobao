package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemDelete 删除单条商品
// taobao.item.delete
//
// 删除单条商品
func TaobaoItemDelete(clt *core.SDKClient, req *tbitem.TaobaoItemDeleteAPIRequest, session string) (*tbitem.TaobaoItemDeleteAPIResponse, error) {
	var resp tbitem.TaobaoItemDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
