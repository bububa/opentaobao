package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取基本和单据信息 API返回值 
alibaba.alihealth.drug.kyt.getcodebillinfo

根据码信息获取基本信息和单据信息
*/
type AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel
}

// 根据码获取基本和单据信息 成功返回结果
type AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getcodebillinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
