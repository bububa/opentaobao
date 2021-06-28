package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单结果同步 APIResponse
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息
*/
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse
}

type TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_identification_case_result_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
