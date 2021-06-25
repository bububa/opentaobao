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
    Response *AlipayBaoxianClaimReturngoodsstatusUpdateResponse `json:"alipay_baoxian_claim_returngoodsstatus_update_response,omitempty"`
}

type AlipayBaoxianClaimReturngoodsstatusUpdateResponse struct {

    // result
    Result   *MtopResult `json:"result,omitempty"`

}
