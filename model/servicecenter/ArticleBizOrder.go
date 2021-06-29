package servicecenter

// ArticleBizOrder 
type ArticleBizOrder struct {
    // 订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 子订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 淘宝会员名
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 应用名称
    ArticleName   string `json:"article_name,omitempty" xml:"article_name,omitempty"`
    // 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    ArticleCode   string `json:"article_code,omitempty" xml:"article_code,omitempty"`
    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 订单创建时间（订购时间）
    Create   string `json:"create,omitempty" xml:"create,omitempty"`
    // 订购周期  1表示年 ，2表示月，3表示天。
    OrderCycle   string `json:"order_cycle,omitempty" xml:"order_cycle,omitempty"`
    // 订购周期开始时间
    OrderCycleStart   string `json:"order_cycle_start,omitempty" xml:"order_cycle_start,omitempty"`
    // 订购周期结束时间
    OrderCycleEnd   string `json:"order_cycle_end,omitempty" xml:"order_cycle_end,omitempty"`
    // 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到）
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 原价（单位为分）
    Fee   string `json:"fee,omitempty" xml:"fee,omitempty"`
    // 优惠（单位为分）
    PromFee   string `json:"prom_fee,omitempty" xml:"prom_fee,omitempty"`
    // 退款（单位为分；升级时，系统会将升级前老版本按照剩余订购天数退还剩余金额）
    RefundFee   string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 实付（单位为分）
    TotalPayFee   string `json:"total_pay_fee,omitempty" xml:"total_pay_fee,omitempty"`
    // 商品模型名称
    ArticleItemName   string `json:"article_item_name,omitempty" xml:"article_item_name,omitempty"`
    // activityCode
    ActivityCode   string `json:"activity_code,omitempty" xml:"activity_code,omitempty"`
}
