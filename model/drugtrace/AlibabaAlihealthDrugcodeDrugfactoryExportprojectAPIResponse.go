package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse 导出项目和药品目录 API返回值
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponseModel is 导出项目和药品目录 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse
func GetAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse {
	return poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse.Get().(*AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse 将 AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse(v *AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse.Put(v)
}
