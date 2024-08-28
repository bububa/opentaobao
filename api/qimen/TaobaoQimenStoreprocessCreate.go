package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreprocessCreate 仓内加工单创建接口
// taobao.qimen.storeprocess.create
//
// ERP调用奇门的接口,创建仓内加工单
func TaobaoQimenStoreprocessCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStoreprocessCreateAPIRequest, resp *qimen.TaobaoQimenStoreprocessCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
