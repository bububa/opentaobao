package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员抵扣规则 APIResponse
alibaba.nlife.b2c.member.discountrule.get

获取企业会员抵扣规则
*/
type AlibabaNlifeB2cMemberDiscountruleGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2cMemberDiscountruleGetResponse
}

type AlibabaNlifeB2cMemberDiscountruleGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2c_member_discountrule_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 业务成功与否 true/false
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 错误码，当result为false时设置
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
    // 错误信息，当result为false时设置
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`

    
    // 抵扣规则
    
    DiscountRule   *DiscountRule `json:"discount_rule,omitempty" xml:"discount_rule,omitempty"`

    
    // 结构化的文案
    
    DiscountMemos   []DiscountMemo `json:"discount_memos,omitempty" xml:"discount_memos>discount_memo,omitempty"`
    
    
}
