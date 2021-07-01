package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗，获取生产企业的存储和运输温度 API返回值 
alibaba.alihealth.drug.kyt.dr.getproteminfo

疫苗，获取生产企业的存储和运输温度
*/
type AlibabaAlihealthDrugKytDrGetproteminfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrGetproteminfoAPIResponseModel
}

// 疫苗，获取生产企业的存储和运输温度 成功返回结果
type AlibabaAlihealthDrugKytDrGetproteminfoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getproteminfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthDrugKytDrGetproteminfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
