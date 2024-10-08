package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerCodeActive 上传激活码的文件
// alibaba.alihealth.tracecodeseller.code.active
//
// 上传商品的激活码文件，存到系统中
func AlibabaAlihealthTracecodesellerCodeActive(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeActiveAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerCodeActiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
