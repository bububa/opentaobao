package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDestbillCheckAPIResponse 直调审批 API返回值
// alibaba.alihealth.drug.kyt.destbill.check
//
// 为药企提供直调单据审批操作
type AlibabaAlihealthDrugKytDestbillCheckAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDestbillCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDestbillCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDestbillCheckAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDestbillCheckAPIResponseModel is 直调审批 成功返回结果
type AlibabaAlihealthDrugKytDestbillCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_destbill_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果标识
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
	// 执行结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDestbillCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.ResponseSuccess = false
	m.Model = false
}

var poolAlibabaAlihealthDrugKytDestbillCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDestbillCheckAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDestbillCheckAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDestbillCheckAPIResponse
func GetAlibabaAlihealthDrugKytDestbillCheckAPIResponse() *AlibabaAlihealthDrugKytDestbillCheckAPIResponse {
	return poolAlibabaAlihealthDrugKytDestbillCheckAPIResponse.Get().(*AlibabaAlihealthDrugKytDestbillCheckAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDestbillCheckAPIResponse 将 AlibabaAlihealthDrugKytDestbillCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDestbillCheckAPIResponse(v *AlibabaAlihealthDrugKytDestbillCheckAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDestbillCheckAPIResponse.Put(v)
}
