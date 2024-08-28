package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreprocessConfirm 仓内加工单确认接口
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
func TaobaoQimenStoreprocessConfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStoreprocessConfirmAPIRequest, resp *qimen.TaobaoQimenStoreprocessConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
