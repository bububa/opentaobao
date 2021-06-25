package servicecenter

// CarLeasePlanDo 
type CarLeasePlanDo struct {

    // 发布商品时选的汽车品牌
    Brand   string `json:"brand,omitempty"`

    // 租赁公司名字
    CompanyName   string `json:"company_name,omitempty"`

    // 无需填写
    Ext   string `json:"ext,omitempty"`

    // 首付金额,单位分
    InitialPaymentAmount   int64 `json:"initial_payment_amount,omitempty"`

    // 首付比例10,15,20代表10%,15%,20%
    InitialPaymentRatio   int64 `json:"initial_payment_ratio,omitempty"`

    // 商品id,外面已经传入，里面可选，传了也会被外面的覆盖
    ItemId   int64 `json:"item_id,omitempty"`

    // 租期12,24,36期
    LeaseTerm   int64 `json:"lease_term,omitempty"`

    // 发布商品时选的汽车车系
    Line   string `json:"line,omitempty"`

    // 融资额,单位分
    LoanAmount   int64 `json:"loan_amount,omitempty"`

    // 发布商品时选的汽车型号
    Model   string `json:"model,omitempty"`

    // 月供,单位分
    MonthlyPayment   int64 `json:"monthly_payment,omitempty"`

    // 官方指导价,单位分
    Msrp   int64 `json:"msrp,omitempty"`

    // 违约金,单位分
    Penalty   int64 `json:"penalty,omitempty"`

    // 卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 门店手续费,单位分
    StoreFee   int64 `json:"store_fee,omitempty"`

    // 尾款金额,单位分
    TailPaymentAmount   int64 `json:"tail_payment_amount,omitempty"`

    // 发布商品时选的汽车年款
    Year   string `json:"year,omitempty"`

    // 支持的尾款类型
    TailPaymentType   string `json:"tail_payment_type,omitempty"`

    // 商家这边唯一用来关联方案的
    OutUniqueId   string `json:"out_unique_id,omitempty"`

    // 违约金实际损失，单位分
    ActualPenalty   int64 `json:"actual_penalty,omitempty"`

}
