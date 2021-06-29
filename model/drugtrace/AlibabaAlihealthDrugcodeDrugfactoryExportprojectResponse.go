package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出项目和药品目录 APIResponse
alibaba.alihealth.drugcode.drugfactory.exportproject

导出临床项目及其药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryExportprojectResponse
}

type AlibabaAlihealthDrugcodeDrugfactoryExportprojectResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_exportproject_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
