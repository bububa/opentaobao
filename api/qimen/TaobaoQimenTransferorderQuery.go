package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderQuery 调拨单查询
// taobao.qimen.transferorder.query
//
// 调拨单查询
func TaobaoQimenTransferorderQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderQueryAPIRequest, resp *qimen.TaobaoQimenTransferorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
