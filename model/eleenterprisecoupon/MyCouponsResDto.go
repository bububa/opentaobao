package eleenterprisecoupon

// MyCouponsResDto 
type MyCouponsResDto struct {

    // 是否可选择多张券
    
    GroupMulti   bool `json:"group_multi,omitempty" xml:"group_multi,omitempty"`
    

    // 券的明细数据
    
    Groups   []CartCouponDto `json:"groups,omitempty" xml:"groups,omitempty"`
    

    // 已选择券的总数量
    
    SelectedQuantity   int64 `json:"selected_quantity,omitempty" xml:"selected_quantity,omitempty"`
    

    // 已选择券的总价，包含溢出与裁剪之后的总价值
    
    SelectedTotalPrice   string `json:"selected_total_price,omitempty" xml:"selected_total_price,omitempty"`
    

    // 推荐券
    
    Recommends   []CartCouponDto `json:"recommends,omitempty" xml:"recommends,omitempty"`
    

    // 选出最优券的数量
    
    RecommendsQuantity   int64 `json:"recommends_quantity,omitempty" xml:"recommends_quantity,omitempty"`
    

    // 选出最优券的总价
    
    RecommendsTotalPrice   string `json:"recommends_total_price,omitempty" xml:"recommends_total_price,omitempty"`
    

    // 已选择的券是不是可以有效组合选中
    
    SelectedLegalStatus   bool `json:"selected_legal_status,omitempty" xml:"selected_legal_status,omitempty"`
    

    // 已选择券不可以组合选中时的提示性文案
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 优惠券名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 优惠券金额
    
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 限制条件
    
    Descriptions   []string `json:"descriptions,omitempty" xml:"descriptions>string,omitempty"`
    

    // 优惠券门槛金额
    
    Threshold   string `json:"threshold,omitempty" xml:"threshold,omitempty"`
    

    // 券ID
    
    Sn   string `json:"sn,omitempty" xml:"sn,omitempty"`
    

    // 券名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 券生效日期时间
    
    BeginDateTime   string `json:"begin_date_time,omitempty" xml:"begin_date_time,omitempty"`
    

    // 券失效日期时间
    
    EndDateTime   string `json:"end_date_time,omitempty" xml:"end_date_time,omitempty"`
    

    // 手机号
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 券使用条件
    
    Condition   string `json:"condition,omitempty" xml:"condition,omitempty"`
    

}
