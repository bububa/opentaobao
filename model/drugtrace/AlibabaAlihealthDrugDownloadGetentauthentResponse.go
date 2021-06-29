package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权企业列表 API返回值 
alibaba.alihealth.drug.download.getentauthent

D2D数据落地获取授权企业列表
*/
type AlibabaAlihealthDrugDownloadGetentauthentAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugDownloadGetentauthentResponse
}

// 获取授权企业列表 成功返回结果
type AlibabaAlihealthDrugDownloadGetentauthentResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getentauthent_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthDrugDownloadGetentauthentResult `json:"result,omitempty" xml:"result,omitempty"`
}
