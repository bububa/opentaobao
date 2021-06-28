package alsc

// CustomerIdQueryOpenReq 
type CustomerIdQueryOpenReq struct {

    // 品牌ID 外部品牌id 2选1
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 店铺ID和外部门店ID必须一
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 顾客ID
    
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    

    // 手机号码
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 物理卡号
    
    PhysicalCardId   string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
    

    // 电子卡号
    
    CardId   string `json:"card_id,omitempty" xml:"card_id,omitempty"`
    

    // 外部品牌id
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

    // 外部门店ID，店铺ID和外部门店ID必须一
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 微信openID
    
    WechatOpenId   string `json:"wechat_open_id,omitempty" xml:"wechat_open_id,omitempty"`
    

    // 微信小程序ID
    
    WechatAppId   string `json:"wechat_app_id,omitempty" xml:"wechat_app_id,omitempty"`
    

    // 座机
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 支付宝ID
    
    AlipayId   string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
    

    // 查询选项，默认查询顾客基础信息， CARD：查询顾客名下的卡列表  ,RECHARGE：查询卡下的储值账户信息  ,POINT：查询顾客的积分信息.
    
    Options   []string `json:"options,omitempty" xml:"options>string,omitempty"`
    

}
