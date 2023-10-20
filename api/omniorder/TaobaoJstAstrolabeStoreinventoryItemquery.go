package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaojstastrolabestoreinventoryitemquery 库存查询接口
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
func Taobaojstastrolabestoreinventoryitemquery(clt *core.SDKClient, req *omniorder.TaobaojstastrolabestoreinventoryitemqueryAPIRequest, session string) (*omniorder.TaobaojstastrolabestoreinventoryitemqueryAPIResponse, error) {
	var resp omniorder.TaobaojstastrolabestoreinventoryitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
