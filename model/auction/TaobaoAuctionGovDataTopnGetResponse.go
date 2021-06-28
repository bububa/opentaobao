package auction

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据不同维度，获取排行榜列表 APIResponse
taobao.auction.gov.data.topn.get

根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
*/
type TaobaoAuctionGovDataTopnGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAuctionGovDataTopnGetResponse `json:"auction_gov_data_topn_get_response,omitempty"` 
    TaobaoAuctionGovDataTopnGetResponse
}

/* model for simplify = false
type TaobaoAuctionGovDataTopnGetResponse struct {

    // 法院维度标的统计排行
    
    Ranks  struct {
        CourtsBidStatTopnDto  []CourtsBidStatTopnDto `json:"courts_bid_stat_topn_dto,omitempty"`
    } `json:"ranks,omitempty"`
    

}
*/

type TaobaoAuctionGovDataTopnGetResponse struct {

    // 法院维度标的统计排行
    Ranks   []CourtsBidStatTopnDto `json:"ranks,omitempty"`

}
