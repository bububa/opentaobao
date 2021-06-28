package simba

// InsightWordPriceDistributeDataDTO 
type InsightWordPriceDistributeDataDTO struct {

    // 展现量
    
    Impression   int64 `json:"impression,omitempty" xml:"impression,omitempty"`
    

    // 点击量
    
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    

    // 花费，单位（分）
    
    Cost   int64 `json:"cost,omitempty" xml:"cost,omitempty"`
    

    // 直接成交金额
    
    Directtransaction   int64 `json:"directtransaction,omitempty" xml:"directtransaction,omitempty"`
    

    // 间接成交金额
    
    Indirecttransaction   int64 `json:"indirecttransaction,omitempty" xml:"indirecttransaction,omitempty"`
    

    // 直接成交笔数
    
    Directtransactionshipping   int64 `json:"directtransactionshipping,omitempty" xml:"directtransactionshipping,omitempty"`
    

    // 间接成交笔数
    
    Indirecttransactionshipping   int64 `json:"indirecttransactionshipping,omitempty" xml:"indirecttransactionshipping,omitempty"`
    

    // 宝贝搜藏数
    
    Favitemtotal   int64 `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
    

    // 店铺搜藏数
    
    Favshoptotal   int64 `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
    

    // 总的成交笔数
    
    Transactionshippingtotal   int64 `json:"transactionshippingtotal,omitempty" xml:"transactionshippingtotal,omitempty"`
    

    // 成交总金额
    
    Transactiontotal   int64 `json:"transactiontotal,omitempty" xml:"transactiontotal,omitempty"`
    

    // 总的收藏数，包括宝贝收藏和店铺收藏
    
    Favtotal   int64 `json:"favtotal,omitempty" xml:"favtotal,omitempty"`
    

    // 竞争度
    
    Competition   int64 `json:"competition,omitempty" xml:"competition,omitempty"`
    

    // 点击率
    
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    

    // 平均点击花费
    
    Cpc   string `json:"cpc,omitempty" xml:"cpc,omitempty"`
    

    // 投入产出比
    
    Roi   string `json:"roi,omitempty" xml:"roi,omitempty"`
    

    // 点击转化率
    
    Coverage   string `json:"coverage,omitempty" xml:"coverage,omitempty"`
    

    // 关键词
    
    Bidword   string `json:"bidword,omitempty" xml:"bidword,omitempty"`
    

    // 竞价区间
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

}
