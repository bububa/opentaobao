package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerProductSearch 查询商品api
// alibaba.alihealth.tracecodeseller.product.search
//
// 查询商品api
func AlibabaAlihealthTracecodesellerProductSearch(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerProductSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerProductSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
