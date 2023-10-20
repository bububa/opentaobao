package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponse 奶粉溯源-同步数据 API返回值
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
type AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponseModel
}

// AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponseModel is 奶粉溯源-同步数据 成功返回结果
type AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponseModel struct {
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
