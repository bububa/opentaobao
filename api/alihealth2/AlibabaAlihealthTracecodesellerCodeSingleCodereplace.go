package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerCodeSingleCodereplace 非药单码替换
// alibaba.alihealth.tracecodeseller.code.single.codereplace
//
// 提供非药追溯码单码替换功能
func AlibabaAlihealthTracecodesellerCodeSingleCodereplace(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
