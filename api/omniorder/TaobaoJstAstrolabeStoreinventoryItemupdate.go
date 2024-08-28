package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryItemupdate 库存增量更新接口
// taobao.jst.astrolabe.storeinventory.itemupdate
//
// ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。
func TaobaoJstAstrolabeStoreinventoryItemupdate(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryItemupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
