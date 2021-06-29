package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出所有项目的药物属性和药品信息 APIResponse
alibaba.alihealth.drugcode.drugfactory.exportattribute

导出所有项目的药物属性和药品信息
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryExportattributeResponse
}

type AlibabaAlihealthDrugcodeDrugfactoryExportattributeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportattribute_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
