package promotion

// BenefitSelectorQuery 
/* model for simplify = false
type BenefitSelectorQuery struct {

    // 权益类型
    
    BenefitType   string `json:"benefit_type,omitempty"`
    

    // 权益id
    
    BenefitId   int64 `json:"benefit_id,omitempty"`
    

    // 选择器步骤,步骤1：不需要填写其他参数，获取拥有的权益类型列表；步骤2：需填写权益类型benefit_type字段，获取该权益类型下的权益列表；步骤3：必填benefit_id，获取对应权益的信息
    
    Step   int64 `json:"step,omitempty"`
    

    // 需要过滤的option
    
    ExcludeOptions   int64 `json:"exclude_options,omitempty"`
    

    // 新权益类型id:1支付宝红包,2优酷会员卡,3不定面额红包,4流量钱包,5淘话费,6淘金币
    
    ConfigId   int64 `json:"config_id,omitempty"`
    

}
*/

// BenefitSelectorQuery 
type BenefitSelectorQuery struct {

    // 权益类型
    BenefitType   string `json:"benefit_type,omitempty"`

    // 权益id
    BenefitId   int64 `json:"benefit_id,omitempty"`

    // 选择器步骤,步骤1：不需要填写其他参数，获取拥有的权益类型列表；步骤2：需填写权益类型benefit_type字段，获取该权益类型下的权益列表；步骤3：必填benefit_id，获取对应权益的信息
    Step   int64 `json:"step,omitempty"`

    // 需要过滤的option
    ExcludeOptions   int64 `json:"exclude_options,omitempty"`

    // 新权益类型id:1支付宝红包,2优酷会员卡,3不定面额红包,4流量钱包,5淘话费,6淘金币
    ConfigId   int64 `json:"config_id,omitempty"`

}
