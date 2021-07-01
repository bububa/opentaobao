package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
储值流水详细 API返回值 
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口
*/
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponseModel
}

// 储值流水详细 成功返回结果
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_account_flowdetail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
