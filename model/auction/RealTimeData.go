package auction

// RealTimeData 结构体
type RealTimeData struct {
	// 在线标的件数
	AuctionOnlineNumber int64 `json:"auction_online_number,omitempty" xml:"auction_online_number,omitempty"`
	// 用户出价次数
	UserBidNumber int64 `json:"user_bid_number,omitempty" xml:"user_bid_number,omitempty"`
	// 开拍标的件数
	AuctionStartNumber int64 `json:"auction_start_number,omitempty" xml:"auction_start_number,omitempty"`
	// 发布标的件数
	AuctionPublishCount int64 `json:"auction_publish_count,omitempty" xml:"auction_publish_count,omitempty"`
	// 报名人数
	UserApplyNumber int64 `json:"user_apply_number,omitempty" xml:"user_apply_number,omitempty"`
	// 意向用户数，包含订阅和报名次数，会累加历史数据
	UserIntentionNumber int64 `json:"user_intention_number,omitempty" xml:"user_intention_number,omitempty"`
	// 今天预计成交金额
	TodayPredictGmv int64 `json:"today_predict_gmv,omitempty" xml:"today_predict_gmv,omitempty"`
	// 用户围观次数，会累加历史数据
	UserViewNumber int64 `json:"user_view_number,omitempty" xml:"user_view_number,omitempty"`
	// 结束标的件数
	AuctionEndNumber int64 `json:"auction_end_number,omitempty" xml:"auction_end_number,omitempty"`
}
