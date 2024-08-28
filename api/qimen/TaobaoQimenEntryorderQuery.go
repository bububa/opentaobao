package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderQuery 入库单查询接口
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
func TaobaoQimenEntryorderQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderQueryAPIRequest, resp *qimen.TaobaoQimenEntryorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
