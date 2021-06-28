package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资产核销接口 APIResponse
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产
*/
type AlibabaAlscCrmOpenAssertVerifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_open_assert_verify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"