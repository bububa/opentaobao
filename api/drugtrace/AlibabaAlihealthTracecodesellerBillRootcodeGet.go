package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesellerBillRootcodeGet 获取最外层包装码
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
func AlibabaAlihealthTracecodesellerBillRootcodeGet(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
