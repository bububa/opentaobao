package baoxian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新理赔单退货货物状态 APIResponse
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态
*/
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse struct {
    model.CommonResponse
    AlipayBaoxianClaimReturngoodsstatusUpdateResponse
}

type AlipayBaoxianClaimReturngoodsstatusUpdateResponse struct {
    XMLName xml.Name `xml:"alipay_baoxian_claim_returngoodsstatus_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MtopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
