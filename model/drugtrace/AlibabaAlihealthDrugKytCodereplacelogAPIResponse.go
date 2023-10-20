package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodereplacelogAPIResponse 码替换记录查询 API返回值
// alibaba.alihealth.drug.kyt.codereplacelog
//
// 码替换记录查询
type AlibabaAlihealthDrugKytCodereplacelogAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytCodereplacelogAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodereplacelogAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytCodereplacelogAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytCodereplacelogAPIResponseModel is 码替换记录查询 成功返回结果
type AlibabaAlihealthDrugKytCodereplacelogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codereplacelog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *ToolPage `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功(true 成功 false 失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodereplacelogAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = nil
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytCodereplacelogAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodereplacelogAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytCodereplacelogAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodereplacelogAPIResponse
func GetAlibabaAlihealthDrugKytCodereplacelogAPIResponse() *AlibabaAlihealthDrugKytCodereplacelogAPIResponse {
	return poolAlibabaAlihealthDrugKytCodereplacelogAPIResponse.Get().(*AlibabaAlihealthDrugKytCodereplacelogAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytCodereplacelogAPIResponse 将 AlibabaAlihealthDrugKytCodereplacelogAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodereplacelogAPIResponse(v *AlibabaAlihealthDrugKytCodereplacelogAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodereplacelogAPIResponse.Put(v)
}
