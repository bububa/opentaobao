package drugtrace

import (
	"encoding/xml"

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

// AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel is 导出所有项目的药物属性和药品信息 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportattribute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
