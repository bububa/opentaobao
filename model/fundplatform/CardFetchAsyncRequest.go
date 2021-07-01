package fundplatform

// CardFetchAsyncRequest 
type CardFetchAsyncRequest struct {
    // 是否储值卡同步激活，默认为true
    Active   bool `json:"active,omitempty" xml:"active,omitempty"`
    // 购买主体类型 company 公司|person  个人
    BuyEntityType   string `json:"buy_entity_type,omitempty" xml:"buy_entity_type,omitempty"`
    // 制卡详情
    CardFetchDetails   []CardFetchDetailDto `json:"card_fetch_details,omitempty" xml:"card_fetch_details>card_fetch_detail_dto,omitempty"`
    // 幂等字段，请保证唯一性,不要超过32位
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    // 售卖方式 online 线上|offline 线下
    SaleMode   string `json:"sale_mode,omitempty" xml:"sale_mode,omitempty"`
    // 业务类型，由资金平台分配
    SubizType   int64 `json:"subiz_type,omitempty" xml:"subiz_type,omitempty"`
    // 购买方在淘宝的用户ID
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
