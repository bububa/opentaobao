package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesellerBillRootcodeGet 获取最外层包装码
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
func AlibabaAlihealthTracecodesellerBillRootcodeGet(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
