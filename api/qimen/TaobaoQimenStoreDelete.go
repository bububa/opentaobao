package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreDelete 门店删除接口
// taobao.qimen.store.delete
//
// 商家在ERP等系统中调用该接口，删除线下门店
func TaobaoQimenStoreDelete(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStoreDeleteAPIRequest, resp *qimen.TaobaoQimenStoreDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
