package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest
奶粉溯源-同步数据 API请求
alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data

奶粉溯源-同步数据 */
type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest struct {
	model.Params
	// 奶粉品牌ID
	_entId string
	// 奶粉数据
	_jsonStr string
}

// New
