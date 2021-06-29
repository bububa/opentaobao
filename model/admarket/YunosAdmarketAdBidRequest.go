package admarket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告竞价服务 APIRequest
yunos.admarket.ad.bid

广告竞价服务
*/
type YunosAdmarketAdBidRequest struct {
    model.Params

    // 竞价请求
    bidRequest   *BidRequest 

}

func NewYunosAdmarketAdBidRequest() *YunosAdmarketAdBidRequest{
    return &YunosAdmarketAdBidRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAdmarketAdBidRequest) GetApiMethodName() string {
    return "yunos.admarket.ad.bid"
}

func (r YunosAdmarketAdBidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAdmarketAdBidRequest) SetBidRequest(bidRequest *BidRequest) error {
    r.bidRequest = bidRequest
    r.Set("bid_request", bidRequest)
    return nil
}

func (r YunosAdmarketAdBidRequest) GetBidRequest() *BidRequest {
    return r.bidRequest
}

