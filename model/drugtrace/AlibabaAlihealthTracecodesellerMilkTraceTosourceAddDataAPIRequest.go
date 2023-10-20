package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest 奶粉溯源-同步数据 API请求
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) Reset() {
	r._entId = ""
	r._jsonStr = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntId is EntId Setter
// 奶粉品牌ID
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetEntId() string {
	return r._entId
}

// SetJsonStr is JsonStr Setter
// 奶粉数据
func (r *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) SetJsonStr(_jsonStr string) error {
	r._jsonStr = _jsonStr
	r.Set("json_str", _jsonStr)
	return nil
}

// GetJsonStr JsonStr Getter
func (r AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) GetJsonStr() string {
	return r._jsonStr
}

var poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest
func GetAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest() *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest {
	return poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest.Get().(*AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest 将 AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest(v *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIRequest.Put(v)
}
