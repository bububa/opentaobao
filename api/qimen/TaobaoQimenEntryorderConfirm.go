package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderConfirm 入库单确认接口
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
func TaobaoQimenEntryorderConfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderConfirmAPIRequest, resp *qimen.TaobaoQimenEntryorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
