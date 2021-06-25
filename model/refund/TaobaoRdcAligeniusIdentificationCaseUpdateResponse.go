package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单信息同步 APIResponse
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息
*/
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRdcAligeniusIdentificationCaseUpdateResponse `json:"taobao_rdc_aligenius_identification_case_update_response,omitempty"`
}

type TaobaoRdcAligeniusIdentificationCaseUpdateResponse struct {

    // result
    Result   *TaobaoRdcAligeniusIdentificationCaseUpdateResult `json:"result,omitempty"`

}
