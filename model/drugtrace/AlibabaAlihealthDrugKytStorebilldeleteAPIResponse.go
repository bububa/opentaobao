package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytStorebilldeleteAPIResponse 零售端单据删除 API返回值
// alibaba.alihealth.drug.kyt.storebilldelete
//
// 零售端单据删除
type AlibabaAlihealthDrugKytStorebilldeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytStorebilldeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel is 零售端单据删除 成功返回结果
type AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_storebilldelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果说明
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 调用信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用信息描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytStorebilldeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytStorebilldeleteAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytStorebilldeleteAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytStorebilldeleteAPIResponse
func GetAlibabaAlihealthDrugKytStorebilldeleteAPIResponse() *AlibabaAlihealthDrugKytStorebilldeleteAPIResponse {
	return poolAlibabaAlihealthDrugKytStorebilldeleteAPIResponse.Get().(*AlibabaAlihealthDrugKytStorebilldeleteAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytStorebilldeleteAPIResponse 将 AlibabaAlihealthDrugKytStorebilldeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytStorebilldeleteAPIResponse(v *AlibabaAlihealthDrugKytStorebilldeleteAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytStorebilldeleteAPIResponse.Put(v)
}
