package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesellerMilkTraceTosourceAddData 奶粉溯源-同步数据
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
func AlibabaAlihealthTracecodesellerMilkTraceTosourceAddData(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
