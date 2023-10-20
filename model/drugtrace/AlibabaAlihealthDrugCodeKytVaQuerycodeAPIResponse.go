package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse 根据码查询码信息 API返回值
// alibaba.alihealth.drug.code.kyt.va.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponseModel is 根据码查询码信息 成功返回结果
type AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_va_querycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugCodeKytVaQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse
func GetAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse() *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse 将 AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse(v *AlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytVaQuerycodeAPIResponse.Put(v)
}
