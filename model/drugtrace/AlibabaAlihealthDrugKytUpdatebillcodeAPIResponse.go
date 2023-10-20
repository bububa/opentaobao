package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse 零售修改出入库单追溯码 API返回值
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
type AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUpdatebillcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUpdatebillcodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUpdatebillcodeAPIResponseModel is 零售修改出入库单追溯码 成功返回结果
type AlibabaAlihealthDrugKytUpdatebillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_updatebillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回接口
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpdatebillcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse
func GetAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse() *AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse {
	return poolAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse.Get().(*AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse 将 AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse(v *AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUpdatebillcodeAPIResponse.Put(v)
}
