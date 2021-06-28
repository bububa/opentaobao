package scbp

// KeywordDto 
/* model for simplify = false
type KeywordDto struct {

    // 效果数据
    
    Effect  *struct {
        KeywordEffectDto  *KeywordEffectDto `json:"keyword_effect_dto,omitempty"`
    } `json:"effect,omitempty"`
    

    // 主键
    
    Id   int64 `json:"id,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 更新时间
    
    GmtModify   string `json:"gmt_modify,omitempty"`
    

    // 产品id
    
    ProductId   int64 `json:"product_id,omitempty"`
    

    // 计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty"`
    

    // 词
    
    Word   string `json:"word,omitempty"`
    

    // 普通词
    
    NormWord   string `json:"norm_word,omitempty"`
    

    // 状态
    
    OnlineStatus   int64 `json:"online_status,omitempty"`
    

    // 出价分
    
    BidPrice   string `json:"bid_price,omitempty"`
    

    // 平均出价
    
    AvgPrice   string `json:"avg_price,omitempty"`
    

    // 建议出价
    
    SugPrice   string `json:"sug_price,omitempty"`
    

    // 低价
    
    BasePrice   string `json:"base_price,omitempty"`
    

    // 星级
    
    QsStar   int64 `json:"qs_star,omitempty"`
    

    // 最优商品id
    
    BestMatchProduct   int64 `json:"best_match_product,omitempty"`
    

    // 相关性产品数量
    
    RelativeProductsCount   int64 `json:"relative_products_count,omitempty"`
    

    // 搜索数量
    
    SearchCount   int64 `json:"search_count,omitempty"`
    

    // 购买数量
    
    BuyCount   int64 `json:"buy_count,omitempty"`
    

    // 配置信息
    
    Properties   string `json:"properties,omitempty"`
    

    // 战略数据
    
    BidStrategy  *struct {
        BidStrategyDto  *BidStrategyDto `json:"bid_strategy_dto,omitempty"`
    } `json:"bid_strategy,omitempty"`
    

}
*/

// KeywordDto 
type KeywordDto struct {

    // 效果数据
    Effect   *KeywordEffectDto `json:"effect,omitempty"`

    // 主键
    Id   int64 `json:"id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModify   string `json:"gmt_modify,omitempty"`

    // 产品id
    ProductId   int64 `json:"product_id,omitempty"`

    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 词
    Word   string `json:"word,omitempty"`

    // 普通词
    NormWord   string `json:"norm_word,omitempty"`

    // 状态
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // 出价分
    BidPrice   string `json:"bid_price,omitempty"`

    // 平均出价
    AvgPrice   string `json:"avg_price,omitempty"`

    // 建议出价
    SugPrice   string `json:"sug_price,omitempty"`

    // 低价
    BasePrice   string `json:"base_price,omitempty"`

    // 星级
    QsStar   int64 `json:"qs_star,omitempty"`

    // 最优商品id
    BestMatchProduct   int64 `json:"best_match_product,omitempty"`

    // 相关性产品数量
    RelativeProductsCount   int64 `json:"relative_products_count,omitempty"`

    // 搜索数量
    SearchCount   int64 `json:"search_count,omitempty"`

    // 购买数量
    BuyCount   int64 `json:"buy_count,omitempty"`

    // 配置信息
    Properties   string `json:"properties,omitempty"`

    // 战略数据
    BidStrategy   *BidStrategyDto `json:"bid_strategy,omitempty"`

}
