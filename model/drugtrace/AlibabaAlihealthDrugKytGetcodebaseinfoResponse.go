package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码的药品信息查询 APIResponse
alibaba.alihealth.drug.kyt.getcodebaseinfo

提供根据码查询码基本信息接口
*/
type AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytGetcodebaseinfoResponse
}

type AlibabaAlihealthDrugKytGetcodebaseinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getcodebaseinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *CodeFullInfoDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
