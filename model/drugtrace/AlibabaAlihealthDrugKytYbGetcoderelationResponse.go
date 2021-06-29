package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
医保-查询码的所有子码 APIResponse
alibaba.alihealth.drug.kyt.yb.getcoderelation

应用于药店或医院入库环节，通过扫码获取下级码进行入库；
通过码查询所有子码以及包装比例
*/
type AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytYbGetcoderelationResponse
}

type AlibabaAlihealthDrugKytYbGetcoderelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yb_getcoderelation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytYbGetcoderelationResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
