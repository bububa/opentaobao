package xhotelofficial

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住用户资格预校验接口 API返回值 
taobao.xhotel.order.official.precheck

官网信用住用户资格预校验接口是在订单创建之前，根据入住人身份信息对其做预先校验是否具有信用住资格。可以优化用户预定体验，对于无资格的用户在预定前即不可进行信用住的选择。减少在提交预定后预定失败体验。该接口为可选对接接口，商家可根据实际情况自行决定是否对接。

接口使用场景

提交订单前的预定人信用住资格预先校验，卖家可决定是否在搜索，预订页，补全身份信息时进行调用，以便决定信用住是否提供给用户
*/
type TaobaoXhotelOrderOfficialPrecheckAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderOfficialPrecheckAPIResponseModel
}

// 官网信用住用户资格预校验接口 成功返回结果
type TaobaoXhotelOrderOfficialPrecheckAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_order_official_precheck_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否符合信用住条件
    MatchCondition   bool `json:"match_condition,omitempty" xml:"match_condition,omitempty"`
    // 当match_condition=false时该字段有意义,用于说明用户不符合信用住条件的原因。以下两种情况，请不要读取此字段值（1、match_condition=true；2、当match_condition=false并且action=1时候（action=1表示用户未签约信用住））
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 入参信息回传, 用于校验的证件号码
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    // 当match_condition=false时该字段有意义,用于标示当用户不符合条件时,应该进行的下一步动作.   0或者空: 用户没有资格使用信用住。可以读取reson字段查看原因。   1: 表示用户符合资格，但是未签约信用住。可以提示用户进行签约后重试。
    Action   int64 `json:"action,omitempty" xml:"action,omitempty"`
}
