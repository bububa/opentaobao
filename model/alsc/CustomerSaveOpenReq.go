package alsc

// CustomerSaveOpenReq 
type CustomerSaveOpenReq struct {

    // 生日
    
    Birthday   string `json:"birthday,omitempty" xml:"birthday,omitempty"`
    

    // saas品牌id
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 渠道  crm_back(1, "CRM后台"),  pos(2, "POS"),  mobile_shop(3, "个人中心"),  wechat(4, "微信"),  alipay(5, "支付宝"),所有pos端传2
    
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    

    // 性别 0女 1男,2其他
    
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 手机号
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 操作人
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 外部id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 外部类型:  wechat：微信openId  alipay：支付宝
    
    OuterType   string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 请求幂等id
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // saas门店id
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

}
