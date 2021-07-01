package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取运营公司信息 API返回值 
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息
*/
type AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel
}

// 根据园区ID获取运营公司信息 成功返回结果
type AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_campus_core_companycampus_getcombycamid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应结果
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
