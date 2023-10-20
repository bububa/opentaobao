package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse 零售端平台单据更新 API返回值
// alibaba.alihealth.drug.kyt.updatebillinfo
//
// 零售端平台单据更新
type AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUpdatebillinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUpdatebillinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUpdatebillinfoAPIResponseModel is 零售端平台单据更新 成功返回结果
type AlibabaAlihealthDrugKytUpdatebillinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_updatebillinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpdatebillinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse
func GetAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse() *AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse 将 AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse(v *AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUpdatebillinfoAPIResponse.Put(v)
}
