package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerCodeRelationCodeantiactive 解除码的关联关系
// alibaba.alihealth.tracecodeseller.code.relation.codeantiactive
//
// 解除码的关联关系
func AlibabaAlihealthTracecodesellerCodeRelationCodeantiactive(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
