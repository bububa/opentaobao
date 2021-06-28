package baoxian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新赔案 APIResponse
alipay.baoxian.claim.update

更新保险理赔单
*/
type AlipayBaoxianClaimUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alipay_baoxian_claim_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MtopResult `json:"result,omitempty" xml:"