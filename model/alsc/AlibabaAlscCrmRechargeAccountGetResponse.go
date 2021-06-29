package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询储值账户信息 APIResponse
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
type AlibabaAlscCrmRechargeAccountGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeAccountGetResponse
}

type AlibabaAlscCrmRechargeAccountGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_account_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
