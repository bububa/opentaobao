package wdk

// MemberInfoDo 
/* model for simplify = false
type MemberInfoDo struct {

    // 会员卡等级
    
    CardLevel   int64 `json:"card_level,omitempty"`
    

    // 会员卡号
    
    CardNum   string `json:"card_num,omitempty"`
    

    // 会员卡状态  '使用中' or '已挂失' or '已作废' or '已补发卡' or '已退卡' or '已冻结' or '未激活' or '已坏卡登记' or '已销毁' or '已制卡' or '已发卡' or '已核对' or '已回收' or '空卡' or '异常' or '已损卡'
    
    State   string `json:"state,omitempty"`
    

    // 卡有效期截止日期，格式：yyyy-MM-dd HH:mm:ss
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 如果卡长期有效，值为true，为true时，默认不校验endTime
    
    Forever   bool `json:"forever,omitempty"`
    

    // 扩展属性map
    
    MemberAttributes   string `json:"member_attributes,omitempty"`
    

    // 会员卡类型
    
    CardType   string `json:"card_type,omitempty"`
    

}
*/

// MemberInfoDo 
type MemberInfoDo struct {

    // 会员卡等级
    CardLevel   int64 `json:"card_level,omitempty"`

    // 会员卡号
    CardNum   string `json:"card_num,omitempty"`

    // 会员卡状态  '使用中' or '已挂失' or '已作废' or '已补发卡' or '已退卡' or '已冻结' or '未激活' or '已坏卡登记' or '已销毁' or '已制卡' or '已发卡' or '已核对' or '已回收' or '空卡' or '异常' or '已损卡'
    State   string `json:"state,omitempty"`

    // 卡有效期截止日期，格式：yyyy-MM-dd HH:mm:ss
    EndTime   string `json:"end_time,omitempty"`

    // 如果卡长期有效，值为true，为true时，默认不校验endTime
    Forever   bool `json:"forever,omitempty"`

    // 扩展属性map
    MemberAttributes   string `json:"member_attributes,omitempty"`

    // 会员卡类型
    CardType   string `json:"card_type,omitempty"`

}
