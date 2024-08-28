package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreitemQuery 门店关联商品查询接口
// taobao.qimen.storeitem.query
//
// 商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
func TaobaoQimenStoreitemQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStoreitemQueryAPIRequest, resp *qimen.TaobaoQimenStoreitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
