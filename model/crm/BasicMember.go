package crm

// BasicMember 
type BasicMember struct {

    // 会员昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 显示会员的状态，normal正常，blacklist黑名单
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 会员等级，0未非会员，其余对应等级名称取grade_name
    
    Grade   int64 `json:"grade,omitempty" xml:"grade,omitempty"`
    

    // 交易成功的次数
    
    TradeCount   int64 `json:"trade_count,omitempty" xml:"trade_count,omitempty"`
    

    // 交易的金额
    
    TradeAmount   float64 `json:"trade_amount,omitempty" xml:"trade_amount,omitempty"`
    

    // 最后交易的日期
    
    LastTradeTime   string `json:"last_trade_time,omitempty" xml:"last_trade_time,omitempty"`
    

    // 交易关闭的笔数
    
    CloseTradeCount   int64 `json:"close_trade_count,omitempty" xml:"close_trade_count,omitempty"`
    

    // 交易关闭金额
    
    CloseTradeAmount   float64 `json:"close_trade_amount,omitempty" xml:"close_trade_amount,omitempty"`
    

    // 购买的宝贝件数
    
    ItemNum   int64 `json:"item_num,omitempty" xml:"item_num,omitempty"`
    

    // 分组的id集合字符串
    
    GroupIds   string `json:"group_ids,omitempty" xml:"group_ids,omitempty"`
    

    // 关系来源，1交易成功，2未交易成功 ,3 卖家主动吸纳
    
    RelationSource   int64 `json:"relation_source,omitempty" xml:"relation_source,omitempty"`
    

    // 最后一次交易的订单号.注:该字段从2014.4.23之后不再返回.
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 等级名称
    
    GradeName   string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
    

}
