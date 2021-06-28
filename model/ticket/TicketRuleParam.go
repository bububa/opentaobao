package ticket

// TicketRuleParam 
type TicketRuleParam struct {

    // 商户景点编码
    
    OutScenicId   string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
    

    // 阿里旅行景点编码
    
    AliScenicId   int64 `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
    

    // 卖家景点规则编码
    
    OutRuleId   string `json:"out_rule_id,omitempty" xml:"out_rule_id,omitempty"`
    

    // 卖家景点规则名称
    
    OutRuleName   string `json:"out_rule_name,omitempty" xml:"out_rule_name,omitempty"`
    

    // 票种规则类型(0实体票规则；1电子票规则)
    
    RuleType   int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
    

    // 规则状态（0：有效，-1：失效）
    
    RuleStatus   int64 `json:"rule_status,omitempty" xml:"rule_status,omitempty"`
    

    // 退票类型。1-无条件退改， 2-有条件退改， 3-不可退改。
    
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    

    // 退票描述
    
    RefundDesc   string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
    

    // 游客信息设置 1:不需要,2:仅需一位游客信息,3:需要所有游客信息
    
    VisitorRequire   int64 `json:"visitor_require,omitempty" xml:"visitor_require,omitempty"`
    

    // 需要的出游人信息，2:手机号,3:身份证,4:姓名,5:护照号码 6:护照姓名拼音。
    
    VisitorInfos   string `json:"visitor_infos,omitempty" xml:"visitor_infos,omitempty"`
    

    // 出游人 是否限购。1:限购,2:不限购
    
    VisitorLimitAble   int64 `json:"visitor_limit_able,omitempty" xml:"visitor_limit_able,omitempty"`
    

    // 限购模式。mode为1按天, mode为2按周, mode为3按月
    
    VisitorLimitMode   int64 `json:"visitor_limit_mode,omitempty" xml:"visitor_limit_mode,omitempty"`
    

    // 限购数量
    
    VisitorLimitNum   int64 `json:"visitor_limit_num,omitempty" xml:"visitor_limit_num,omitempty"`
    

    // 入园类型。1-用兑换凭证直接入园，2-用兑换凭证换票入园
    
    EnterType   int64 `json:"enter_type,omitempty" xml:"enter_type,omitempty"`
    

    // 入园使用的凭证类型。1、二维码，2、身份证，3、二维码或身份证，4:数字码，5、手机号，6、其它
    
    EnterVoucherType   int64 `json:"enter_voucher_type,omitempty" xml:"enter_voucher_type,omitempty"`
    

    // 其他入园凭证类型。
    
    EnterVoucherValue   string `json:"enter_voucher_value,omitempty" xml:"enter_voucher_value,omitempty"`
    

    // 换票地址
    
    TicketChangeAdderss   string `json:"ticket_change_adderss,omitempty" xml:"ticket_change_adderss,omitempty"`
    

    // 景区入园地址
    
    EnterAddress   string `json:"enter_address,omitempty" xml:"enter_address,omitempty"`
    

    // 门票费用包含
    
    FeeInclude   string `json:"fee_include,omitempty" xml:"fee_include,omitempty"`
    

    // 补充说明
    
    ExtraDesc   string `json:"extra_desc,omitempty" xml:"extra_desc,omitempty"`
    

    // 限购类型。0-身份证限购， 1-手机号限购
    
    VisitorLimitType   int64 `json:"visitor_limit_type,omitempty" xml:"visitor_limit_type,omitempty"`
    

}
