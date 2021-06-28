package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单信息同步 APIResponse
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息
*/
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusIdentificationCaseUpdateResponse
}

type TaobaoRdcAligeniusIdentificationCaseUpdateResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_identification_case_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusIdentificationCaseUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
