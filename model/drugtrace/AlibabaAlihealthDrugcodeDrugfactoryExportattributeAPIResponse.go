package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse 导出所有项目的药物属性和药品信息 API返回值
// alibaba.alihealth.drugcode.drugfactory.exportattribute
//
// 导出所有项目的药物属性和药品信息
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel is 导出所有项目的药物属性和药品信息 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportattribute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse
func GetAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse {
	return poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse.Get().(*AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse 将 AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse(v *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse.Put(v)
}
