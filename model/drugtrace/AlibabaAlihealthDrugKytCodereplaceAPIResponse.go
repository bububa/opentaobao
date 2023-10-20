package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodereplaceAPIResponse 单码替换 API返回值
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
type AlibabaAlihealthDrugKytCodereplaceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytCodereplaceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodereplaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytCodereplaceAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytCodereplaceAPIResponseModel is 单码替换 成功返回结果
type AlibabaAlihealthDrugKytCodereplaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_codereplace_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用是否成功(true 成功 false 失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytCodereplaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytCodereplaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytCodereplaceAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytCodereplaceAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodereplaceAPIResponse
func GetAlibabaAlihealthDrugKytCodereplaceAPIResponse() *AlibabaAlihealthDrugKytCodereplaceAPIResponse {
	return poolAlibabaAlihealthDrugKytCodereplaceAPIResponse.Get().(*AlibabaAlihealthDrugKytCodereplaceAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytCodereplaceAPIResponse 将 AlibabaAlihealthDrugKytCodereplaceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodereplaceAPIResponse(v *AlibabaAlihealthDrugKytCodereplaceAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodereplaceAPIResponse.Put(v)
}
