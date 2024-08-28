package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderCreate 入库单创建接口
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
func TaobaoQimenEntryorderCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderCreateAPIRequest, resp *qimen.TaobaoQimenEntryorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
