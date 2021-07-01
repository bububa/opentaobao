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

// NewAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest 初始化AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest对象
func NewAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest() *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest {
	return &AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EntId Setter
// 奶粉品牌ID
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// Get EntId Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetEntId() string {
	return r._entId
}

// Set is JsonStr Setter
// 奶粉数据
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) SetJsonStr(_jsonStr string) error {
	r._jsonStr = _jsonStr
	r.Set("json_str", _jsonStr)
	return nil
}

// Get JsonStr Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetJsonStr() string {
	return r._jsonStr
}
