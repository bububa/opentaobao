package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
企业下载列表 APIResponse
alibaba.alihealth.drug.download.entlist

获取企业的下载文件列表
*/
type AlibabaAlihealthDrugDownloadEntlistAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugDownloadEntlistResponse
}

type AlibabaAlihealthDrugDownloadEntlistResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_download_entlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAlihealthDrugDownloadEntlistResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
