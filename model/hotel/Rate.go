package hotel

// Rate 
type Rate struct {

    // 餐食信息
    Breakfast   string `json:"breakfast,omitempty"`

    // 餐食类型，0--无早, 1--单早, 2--双早, 3--三早, 4--多早
    BreakfastCount   int64 `json:"breakfast_count,omitempty"`

    // 最晚入住时间
    CanCheckinEnd   string `json:"can_checkin_end,omitempty"`

    // 最早入住时间
    CanCheckinStart   string `json:"can_checkin_start,omitempty"`

    // 退订政策描述，默认是短文案；
    CancelPolicyDesc   string `json:"cancel_policy_desc,omitempty"`

    // 退订政策长文案
    CancelPolicyDescLong   string `json:"cancel_policy_desc_long,omitempty"`

    // 退订政策中文案
    CancelPolicyDescMiddle   string `json:"cancel_policy_desc_middle,omitempty"`

    // 退订政策类型枚举，1-不能退, 2-免费退, 3-按规则退
    CancelType   int64 `json:"cancel_type,omitempty"`

    // 担保开始时间
    GuaranteeStartTime   string `json:"guarantee_start_time,omitempty"`

    // 担保类型
    GuaranteeType   int64 `json:"guarantee_type,omitempty"`

    // h5下单页链接
    H5BuyUrl   string `json:"h5_buy_url,omitempty"`

    // 最长可入住时间标识，float类型
    Hourage   string `json:"hourage,omitempty"`

    // 是否立即确认
    InstantConfirm   bool `json:"instant_confirm,omitempty"`

    // 库存价格信息
    InventoryPrice   string `json:"inventory_price,omitempty"`

    // 会员等级
    MemberLevel   int64 `json:"member_level,omitempty"`

    // 会员等级名称
    MemberLevelName   string `json:"member_level_name,omitempty"`

    // 报价类型
    PaymentTypeByte   int64 `json:"payment_type_byte,omitempty"`

    // pc购买地址
    PcBuyUrl   string `json:"pc_buy_url,omitempty"`

    // 报价id
    RateId   int64 `json:"rate_id,omitempty"`

    // 价格政策名称
    RatePlanName   string `json:"rate_plan_name,omitempty"`

    // 注册状态，true-已注册，false-未注册
    RegisterStatus   bool `json:"register_status,omitempty"`

    // 价格政策id
    RpId   int64 `json:"rp_id,omitempty"`

    // 立减金额； 最新接口数据已包含UMP优惠信息，用于信用住场景的优惠返回；
    Subtract   int64 `json:"subtract,omitempty"`

    // 是否消失方，true--是，空或false-不是
    HourRate   bool `json:"hour_rate,omitempty"`

    // 当前报价对应的rid
    Rid   int64 `json:"rid,omitempty"`

    // 当前报价所属卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 当前报价对应的srid
    Srid   int64 `json:"srid,omitempty"`

    // 最晚离店时间，HH:MM格式
    CanCheckOutEnd   string `json:"can_check_out_end,omitempty"`

    // 是否展示”信用住“标签，true--展示，false
    LaterPay   bool `json:"later_pay,omitempty"`

    // 是否展示“会员价”标签；不是原价的价格上也有优惠价与会员价的区分；true--展示
    MemberPrice   bool `json:"member_price,omitempty"`

    // 卖家供应商渠道代号，一个卖家可以有多个供应商，使用代号进行区别
    Supplier   string `json:"supplier,omitempty"`

}
