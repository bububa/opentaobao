package alihealth2

// BaseRule 
type BaseRule struct {

    // 级别规则
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 规则的详细数据
    
    Detail   string `json:"detail,omitempty" xml:"detail,omitempty"`
    

    // 提示文案
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 周期限购：CycleLimitRule，单次限购规则：SingleLimitRule，频次规则：FreqDosageRule， 每日剂量规则：DailyDosageRule，单次剂量规则：SingleDosageRule ，用法规则：WayDosageRule ，重复用药规则：RepetitionRule ，特殊年龄规则：AgeSpecialCrowdRule ，人群标签规则：UserLabelSpecialCrowdRule ，性别诊断冲突规则则：GenderSpecialCrowdRule ，适应症规则：IndicationRule ，相互作用规则：InteractionRule ，禁忌症规则：ContraIndicationRule
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // spuid
    
    SpuId   string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
    

    // 每次剂量上线（命中每日剂量规则+单次剂量规则）
    
    Threshold   string `json:"threshold,omitempty" xml:"threshold,omitempty"`
    

    // 建议的每日使用频次，如2-3（命中频次规则）
    
    SuggestedFreq   string `json:"suggested_freq,omitempty" xml:"suggested_freq,omitempty"`
    

    // 建议的用法， 如：口服（命中用法规则）
    
    SuggestedWay   string `json:"suggested_way,omitempty" xml:"suggested_way,omitempty"`
    

    // 药品通用名
    
    CommonName   string `json:"common_name,omitempty" xml:"common_name,omitempty"`
    

    // 药品通用名（重复用药+相互作用规则）
    
    OtherCommonName   string `json:"other_common_name,omitempty" xml:"other_common_name,omitempty"`
    

    // 禁用的年龄区间，单位是天，比如0-1000（命中特殊年龄规则）
    
    ForbiddenAge   string `json:"forbidden_age,omitempty" xml:"forbidden_age,omitempty"`
    

    // 禁用的人群（命中特殊人群规则）
    
    ForbiddenUser   string `json:"forbidden_user,omitempty" xml:"forbidden_user,omitempty"`
    

    // 性别（命中性别诊断冲突规则）
    
    Gender   string `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 诊断码（命中性别诊断冲突规则）
    
    DiagCode   string `json:"diag_code,omitempty" xml:"diag_code,omitempty"`
    

    // 诊断（命中适应症规则）
    
    Diags   string `json:"diags,omitempty" xml:"diags,omitempty"`
    

    // 跟该药冲突的诊断码， 如N19.x00（（命中禁忌症规则））
    
    ForbiddenDiags   string `json:"forbidden_diags,omitempty" xml:"forbidden_diags,omitempty"`
    

    // 限购上限制（命中周期限购+单次限购规则）
    
    LimitThreshold   int64 `json:"limit_threshold,omitempty" xml:"limit_threshold,omitempty"`
    

}
