package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemOperateDelete 商品删除
// alibaba.item.operate.delete
//
// 商品删除
func AlibabaItemOperateDelete(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemOperateDeleteAPIRequest, resp *tbitem.AlibabaItemOperateDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
