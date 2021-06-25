package baoxian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新赔案 APIResponse
alipay.baoxian.claim.update

更新保险理赔单
*/
type AlipayBaoxianClaimUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlipayBaoxianClaimUpdateResponse `json:"alipay_baoxian_claim_update_response,omitempty"`
}

type AlipayBaoxianClaimUpdateResponse struct {

    // result
    Result   *MtopResult `json:"result,omitempty"`

}
