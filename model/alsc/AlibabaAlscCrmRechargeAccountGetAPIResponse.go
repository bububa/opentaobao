package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询储值账户信息 API返回值 
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
type AlibabaAlscCrmRechargeAccountGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeAccountGetAPIResponseModel
}

// 查询储值账户信息 成功返回结果
type AlibabaAlscCrmRechargeAccountGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_account_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
