package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemEditSubmit 商品编辑提交schema信息
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
func AlibabaItemEditSubmit(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemEditSubmitAPIRequest, resp *tbitem.AlibabaItemEditSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
