package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
扫码识别会员接口 APIResponse
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
*/
type AlibabaWdkMemberQrcodeIdentifyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMemberQrcodeIdentifyResponse `json:"alibaba_wdk_member_qrcode_identify_response,omitempty"` 
    AlibabaWdkMemberQrcodeIdentifyResponse
}

/* model for simplify = false
type AlibabaWdkMemberQrcodeIdentifyResponse struct {

    // result
    
    Result  *struct {
        MtopResult  *MtopResult `json:"mtop_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMemberQrcodeIdentifyResponse struct {

    // result
    Result   *MtopResult `json:"result,omitempty"`

}
