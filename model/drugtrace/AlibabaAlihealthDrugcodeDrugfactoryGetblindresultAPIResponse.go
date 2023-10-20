package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse 获取盲底文件处理结果 API返回值
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel is 获取盲底文件处理结果 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_getblindresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse
func GetAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse {
	return poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse.Get().(*AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse 将 AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse(v *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse.Put(v)
}
