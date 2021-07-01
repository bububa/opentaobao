package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthTracecodesellerMilkTraceTosourceAddData
奶粉溯源-同步数据
alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data

奶粉溯源-同步数据 */
func AlibabaAlihealthTracecodesellerMilkTraceTosourceAddData(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest, session string) (*drugtrace.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
