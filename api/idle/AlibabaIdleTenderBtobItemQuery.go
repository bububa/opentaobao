package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemQuery 暗拍b2b商品查询
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
func AlibabaIdleTenderBtobItemQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemQueryAPIRequest, resp *idle.AlibabaIdleTenderBtobItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
