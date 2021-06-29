package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取运营公司信息 APIResponse
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息
*/
type AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusCoreCompanycampusGetcombycamidResponse
}

type AlibabaCampusCoreCompanycampusGetcombycamidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_core_companycampus_getcombycamid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应结果
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
