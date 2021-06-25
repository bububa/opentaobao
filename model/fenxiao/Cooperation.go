package fenxiao

// Cooperation 
type Cooperation struct {

    // 合作关系ID
    CooperateId   int64 `json:"cooperate_id,omitempty"`

    // 分销商ID
    DistributorId   int64 `json:"distributor_id,omitempty"`

    // 分销商nick
    DistributorNick   string `json:"distributor_nick,omitempty"`

    // 授权产品线
    ProductLine   string `json:"product_line,omitempty"`

    // 等级ID
    GradeId   int64 `json:"grade_id,omitempty"`

    // 分销方式： AGENT(代销) 、DEALER(经销)
    TradeType   string `json:"trade_type,omitempty"`

    // 供应商授权的支付方式：ALIPAY(支付宝)、OFFPREPAY(预付款)、OFFTRANSFER(转帐)、OFFSETTLEMENT(后期结算)
    AuthPayway   []String `json:"auth_payway,omitempty"`

    // 供应商ID
    SupplierId   int64 `json:"supplier_id,omitempty"`

    // 供应商NICK
    SupplierNick   string `json:"supplier_nick,omitempty"`

    // 合作起始时间
    StartDate   string `json:"start_date,omitempty"`

    // 合作终止时间
    EndDate   string `json:"end_date,omitempty"`

    // 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
    Status   string `json:"status,omitempty"`

    // 授权产品线名称，和product_line中的值按序对应
    ProductLineName   []String `json:"product_line_name,omitempty"`

}
