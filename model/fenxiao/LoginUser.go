package fenxiao

// LoginUser 
type LoginUser struct {

    // 分销用户ID
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 会员NICK
    
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    

    // 分销用户类型(1:分销商，2:供应商，3:品牌商；2、3都表示有供货能力，只是身份不同)
    
    UserType   string `json:"user_type,omitempty" xml:"user_type,omitempty"`
    

    // 入驻时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

}
