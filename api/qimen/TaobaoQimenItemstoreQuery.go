package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenItemstoreQuery 商品关联门店查询接口
// taobao.qimen.itemstore.query
//
// 商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
func TaobaoQimenItemstoreQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenItemstoreQueryAPIRequest, resp *qimen.TaobaoQimenItemstoreQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
