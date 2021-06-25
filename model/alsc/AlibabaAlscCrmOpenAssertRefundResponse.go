package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资产核销回退接口 APIResponse
alibaba.alsc.crm.open.assert.refund

回退已经核销的储值积分券资产
*/
type AlibabaAlscCrmOpenAssertRefundAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmOpenAssertRefundResponse `json:"alibaba_alsc_crm_open_assert_refund_response,omitempty"`
}

type AlibabaAlscCrmOpenAssertRefundResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
