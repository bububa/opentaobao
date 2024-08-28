package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerEntSearch 查询商家信息
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
func AlibabaAlihealthTracecodesellerEntSearch(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
