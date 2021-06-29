package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付链路中的积分抵扣是否合法 API返回值 
alibaba.alsc.crm.point.chkpntbypay

校验支付链路中的积分抵扣是否合法
*/
type AlibabaAlscCrmPointChkpntbypayAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmPointChkpntbypayResponse
}

// 校验支付链路中的积分抵扣是否合法 成功返回结果
type AlibabaAlscCrmPointChkpntbypayResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_point_chkpntbypay_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
