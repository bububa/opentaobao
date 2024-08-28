package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreCreate 门店新增接口
// taobao.qimen.store.create
//
// isv调用接口来讲线下门店同步到线上
func TaobaoQimenStoreCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStoreCreateAPIRequest, resp *qimen.TaobaoQimenStoreCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
