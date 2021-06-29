package traveltrade

// SellerInfo 
type SellerInfo struct {

    // 卖家是否可以对订单进行评价
    
    SellerCanRate   bool `json:"seller_can_rate,omitempty" xml:"seller_can_rate,omitempty"`
    

    // 卖家邮件地址
    
    SellerEmail   string `json:"seller_email,omitempty" xml:"seller_email,omitempty"`
    

    // 卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）红、黄、绿、蓝、紫 分别对应 1、2、3、4、5
    
    SellerFlag   int64 `json:"seller_flag,omitempty" xml:"seller_flag,omitempty"`
    

    // 卖家备注
    
    SellerMemo   string `json:"seller_memo,omitempty" xml:"seller_memo,omitempty"`
    

    // 卖家姓名
    
    SellerName   string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 卖家联系方式
    
    SellerPhone   string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
    

    // 卖家是否已评价。可选值:true(已评价),false(未评价)
    
    SellerRate   bool `json:"seller_rate,omitempty" xml:"seller_rate,omitempty"`
    

    // 卖家店铺名称
    
    SellerShop   string `json:"seller_shop,omitempty" xml:"seller_shop,omitempty"`
    

}
