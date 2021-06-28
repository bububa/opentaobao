package wdk

// FullRangeActivity 
/* model for simplify = false
type FullRangeActivity struct {

    // 通用限购信息，-1为不限制，默认为不限制
    
    LimitInfo  *struct {
        LimitInfo  *LimitInfo `json:"limit_info,omitempty"`
    } `json:"limit_info,omitempty"`
    

    // 优惠适用场景[APP|POS|POS+APP分别对应的值为1|2|1,2]
    
    Terminals  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"terminals,omitempty"`
    

    // 商家活动id
    
    OutActId   string `json:"out_act_id,omitempty"`
    

    // 参加活动的渠道店ids
    
    ShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"shop_ids,omitempty"`
    

    // 活动结束时间，时间戳
    
    EndTime   int64 `json:"end_time,omitempty"`
    

    // 活动的梯度列表
    
    RuleStairs  struct {
        Rulestairs  []Rulestairs `json:"rulestairs,omitempty"`
    } `json:"rule_stairs,omitempty"`
    

    // 活动开始时间，时间戳
    
    StartTime   int64 `json:"start_time,omitempty"`
    

    // 活动详情描述,不超过30个英文字符
    
    Description   string `json:"description,omitempty"`
    

    // 活动名称,不超过10个英文字符
    
    ActivityName   string `json:"activity_name,omitempty"`
    

    // 商家人群编码
    
    MerchantCrowdCode   string `json:"merchant_crowd_code,omitempty"`
    

    // 淘鲜达人群编码
    
    TxdCrowdCode   string `json:"txd_crowd_code,omitempty"`
    

}
*/

// FullRangeActivity 
type FullRangeActivity struct {

    // 通用限购信息，-1为不限制，默认为不限制
    LimitInfo   *LimitInfo `json:"limit_info,omitempty"`

    // 优惠适用场景[APP|POS|POS+APP分别对应的值为1|2|1,2]
    Terminals   []int64 `json:"terminals,omitempty"`

    // 商家活动id
    OutActId   string `json:"out_act_id,omitempty"`

    // 参加活动的渠道店ids
    ShopIds   []string `json:"shop_ids,omitempty"`

    // 活动结束时间，时间戳
    EndTime   int64 `json:"end_time,omitempty"`

    // 活动的梯度列表
    RuleStairs   []Rulestairs `json:"rule_stairs,omitempty"`

    // 活动开始时间，时间戳
    StartTime   int64 `json:"start_time,omitempty"`

    // 活动详情描述,不超过30个英文字符
    Description   string `json:"description,omitempty"`

    // 活动名称,不超过10个英文字符
    ActivityName   string `json:"activity_name,omitempty"`

    // 商家人群编码
    MerchantCrowdCode   string `json:"merchant_crowd_code,omitempty"`

    // 淘鲜达人群编码
    TxdCrowdCode   string `json:"txd_crowd_code,omitempty"`

}
