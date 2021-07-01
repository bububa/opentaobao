package wdk

// DrfTxdActivitySkuBo 结构体
type DrfTxdActivitySkuBo struct {
	// 商品对应的活动Id，仅当同一次任务有相关活动更新的时候在传入
	ActivityVersionId int64 `json:"activity_version_id,omitempty" xml:"activity_version_id,omitempty"`
	// 限购权重
	LimitWeight int64 `json:"limit_weight,omitempty" xml:"limit_weight,omitempty"`
	// 更新时间
	UpdateTime int64 `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 插入时间
	InsertTime int64 `json:"insert_time,omitempty" xml:"insert_time,omitempty"`
	// 状态：0--不可用，1--可用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 门槛数量：件/金额（分）
	ConditionNum int64 `json:"condition_num,omitempty" xml:"condition_num,omitempty"`
	// 门槛类型：2-累计金额消费，3-累计购买次数消费
	ConditionType int64 `json:"condition_type,omitempty" xml:"condition_type,omitempty"`
	// 赠品skuCode
	GiftSkuCode string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
	// 买赠门槛
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
	// 活动每日限购
	TotalDayLimit int64 `json:"total_day_limit,omitempty" xml:"total_day_limit,omitempty"`
	// 用户每日限购
	UserDayLimit int64 `json:"user_day_limit,omitempty" xml:"user_day_limit,omitempty"`
	// 总限购数量
	TotalLimit int64 `json:"total_limit,omitempty" xml:"total_limit,omitempty"`
	// 用户限购
	UserLimit int64 `json:"user_limit,omitempty" xml:"user_limit,omitempty"`
	// 打折
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 减钱
	DecreaseMoney int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
	// 一口价
	FixPrice int64 `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品池ID
	PoolId int64 `json:"pool_id,omitempty" xml:"pool_id,omitempty"`
	// 所属活动ID
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 大润发活动类型
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 淘鲜达活动Id
	TxdActivityId int64 `json:"txd_activity_id,omitempty" xml:"txd_activity_id,omitempty"`
}
