package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资产核销接口 APIResponse
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产
*/
type AlibabaAlscCrmOpenAssertVerifyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmOpenAssertVerifyResponse `json:"alibaba_alsc_crm_open_assert_verify_response,omitempty"` 
    AlibabaAlscCrmOpenAssertVerifyResponse
}

/* model for simplify = false
type AlibabaAlscCrmOpenAssertVerifyResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmOpenAssertVerifyResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
