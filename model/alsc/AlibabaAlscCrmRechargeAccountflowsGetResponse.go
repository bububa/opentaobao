package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询储值流水 API返回值 
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
type AlibabaAlscCrmRechargeAccountflowsGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeAccountflowsGetResponse
}

// 分页查询储值流水 成功返回结果
type AlibabaAlscCrmRechargeAccountflowsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_accountflows_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
