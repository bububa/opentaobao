package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse 奶粉溯源-同步数据 API返回值
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponseModel is 奶粉溯源-同步数据 成功返回结果
type AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_milk_trace_tosource_add_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 服务出参true
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
}

var poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse
func GetAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse() *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse {
	return poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse.Get().(*AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse 将 AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse(v *AlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerMilkTraceTosourceAddDataAPIResponse.Put(v)
}
