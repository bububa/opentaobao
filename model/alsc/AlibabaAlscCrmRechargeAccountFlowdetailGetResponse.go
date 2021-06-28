package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值流水详细 APIResponse
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口
*/
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeAccountFlowdetailGetResponse
}

type AlibabaAlscCrmRechargeAccountFlowdetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_account_flowdetail_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
