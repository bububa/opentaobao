package drug

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrSpuQuery 获取标品库标品信息
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
func AlibabaAlihealthNrSpuQuery(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaAlihealthNrSpuQueryAPIRequest, resp *drug.AlibabaAlihealthNrSpuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
