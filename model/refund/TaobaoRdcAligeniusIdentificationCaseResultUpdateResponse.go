package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单结果同步 APIResponse
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息
*/
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse `json:"rdc_aligenius_identification_case_result_update_response,omitempty"` 
    TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse
}

/* model for simplify = false
type TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoRdcAligeniusIdentificationCaseResultUpdateResult  *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult `json:"taobao_rdc_aligenius_identification_case_result_update_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRdcAligeniusIdentificationCaseResultUpdateResponse struct {

    // 接口返回model
    Result   *TaobaoRdcAligeniusIdentificationCaseResultUpdateResult `json:"result,omitempty"`

}
