package auction

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取司法拍卖最新出价数据 APIResponse
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据
*/
type TaobaoAuctionGovGetLatestbidAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAuctionGovGetLatestbidResponse `json:"taobao_auction_gov_get_latestbid_response,omitempty"`
}

type TaobaoAuctionGovGetLatestbidResponse struct {

    // result
    Result   *Result4Top `json:"result,omitempty"`

}
