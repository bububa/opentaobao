package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取盲底文件处理结果 APIResponse
alibaba.alihealth.drugcode.drugfactory.getblindresult

获取盲底文件处理结果
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryGetblindresultResponse
}

type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_getblindresult_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
