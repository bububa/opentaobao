package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytIdgenerateAPIResponse 终端(医疗机构|零售药店)ID生成接口 API返回值
// alibaba.alihealth.drug.kyt.idgenerate
//
// 终端(医疗机构|零售药店)ID生成接口
type AlibabaAlihealthDrugKytIdgenerateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytIdgenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytIdgenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytIdgenerateAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytIdgenerateAPIResponseModel is 终端(医疗机构|零售药店)ID生成接口 成功返回结果
type AlibabaAlihealthDrugKytIdgenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_idgenerate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的ID
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回的编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回的结果
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回的结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytIdgenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytIdgenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytIdgenerateAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytIdgenerateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytIdgenerateAPIResponse
func GetAlibabaAlihealthDrugKytIdgenerateAPIResponse() *AlibabaAlihealthDrugKytIdgenerateAPIResponse {
	return poolAlibabaAlihealthDrugKytIdgenerateAPIResponse.Get().(*AlibabaAlihealthDrugKytIdgenerateAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytIdgenerateAPIResponse 将 AlibabaAlihealthDrugKytIdgenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytIdgenerateAPIResponse(v *AlibabaAlihealthDrugKytIdgenerateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytIdgenerateAPIResponse.Put(v)
}
