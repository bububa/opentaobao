package crm

// GradePromotion 
type GradePromotion struct {
    // 买家会员级别  0:店铺客户  1：普通会员 2：高级会员 3：VIP会员 4：至尊VIP
    CurGrade   string `json:"cur_grade,omitempty" xml:"cur_grade,omitempty"`
    // 店铺客户、普通会员 、高级会员、VIP会员、至尊VIP
    CurGradeName   string `json:"cur_grade_name,omitempty" xml:"cur_grade_name,omitempty"`
    // 会员级别折扣率没有小数，990代表9.9折
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 升级到下一个级别的需要的交易额，单位：分
    NextUpgradeAmount   int64 `json:"next_upgrade_amount,omitempty" xml:"next_upgrade_amount,omitempty"`
    // 升级到下一个级别的需要的交易量
    NextUpgradeCount   int64 `json:"next_upgrade_count,omitempty" xml:"next_upgrade_count,omitempty"`
    // 该等级对应的下一等级1:普通会员  2：高级会员 3：VIP会员 4：至尊VIP 0：后续等级都未启用或没有下一等级
    NextGradeName   string `json:"next_grade_name,omitempty" xml:"next_grade_name,omitempty"`
    // 普通会员、高级会员、VIP会员、至尊VIP。空的时候代表后续等级未启用或没有下一等级
    NextGrade   string `json:"next_grade,omitempty" xml:"next_grade,omitempty"`
}
