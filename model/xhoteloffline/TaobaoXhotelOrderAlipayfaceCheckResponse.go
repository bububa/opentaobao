package xhoteloffline

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住买家资格校验接口 API返回值 
taobao.xhotel.order.alipayface.check

接口用于校验买家是否具有使用酒店信用住的资格
*/
type TaobaoXhotelOrderAlipayfaceCheckAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceCheckResponse
}

// 线下信用住买家资格校验接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceCheckResponse struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 当match_condition=false时该字段有意义,用于说明用户不符合信用住条件的原因, 主要有如下几种返回:1. 该用户尚未签约线下信用住, 请先扫码签约; 2.该用户尚未签约, 暂不支持使用线下信用住, 请使用现金或其他方式结账; 3. 该用户支付宝账号存在风险, 暂不支持使用线下信用住, 请使用现金或其他方式结账; 4.该用户当前信用额度不足, 无法支付此房费, 请使用现金或其他方式结账
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 当match_condition=false时该字段有意义,用于标示当用户不符合条件时,应该进行的下一步动作. 0或者空: 代表没有下一步动作, 接入方此时可以直接不提示用户走任何线下信用住的流程和文案等; 1: 提示用户进行扫码签约,此时可以把reason字段展示到前台,或者自己定义提示文案
    Action   int64 `json:"action,omitempty" xml:"action,omitempty"`
    // 是否符合信用住条件
    MatchCondition   bool `json:"match_condition,omitempty" xml:"match_condition,omitempty"`
    // 入参信息回传, 用于校验的证件号码
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    // existAlipayOrder
    ExistAlipayOrder   bool `json:"exist_alipay_order,omitempty" xml:"exist_alipay_order,omitempty"`
}
