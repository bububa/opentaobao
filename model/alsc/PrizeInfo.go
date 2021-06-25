package alsc

// PrizeInfo 
type PrizeInfo struct {

    // 优惠券列表
    PrizeItemInfoList   []PrizeItemInfo `json:"prize_item_info_list,omitempty"`

    // 是否中奖
    WinPrize   bool `json:"win_prize,omitempty"`

}
