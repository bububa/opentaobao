package baoxian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新理赔单退货货物状态 APIResponse
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态
*/
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlipayBaoxianClaimReturngoodsstatusUpdateResponse `json:"alipay_baoxian_claim_returngoodsstatus_update_response,omitempty"` 
    AlipayBaoxianClaimReturngoodsstatusUpdateResponse
}

/* model for simplify = false
type AlipayBaoxianClaimReturngoodsstatusUpdateResponse struct {

    // result
    
    Result  *struct {
        MtopResult  *MtopResult `json:"mtop_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlipayBaoxianClaimReturngoodsstatusUpdateResponse struct {

    // result
    Result   *MtopResult `json:"result,omitempty"`

}
