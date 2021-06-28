package user

// UserCredit 
type UserCredit struct {

    // 信用等级（是根据score生成的），信用等级：淘宝会员在淘宝网上的信用度，分为20个级别，级别如：level = 1 时，表示一心；level = 2 时，表示二心
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 信用总分（“好评”加一分，“中评”不加分，“差评”扣一分。分越高，等级越高）
    
    Score   int64 `json:"score,omitempty" xml:"score,omitempty"`
    

    // 收到的评价总条数。取值范围:大于零的整数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    

    // 收到的好评总条数。取值范围:大于零的整数
    
    GoodNum   int64 `json:"good_num,omitempty" xml:"good_num,omitempty"`
    

}
