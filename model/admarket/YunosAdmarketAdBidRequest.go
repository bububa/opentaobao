package admarket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告竞价服务 API请求
yunos.admarket.ad.bid

广告竞价服务
*/
type YunosAdmarketAdBidRequest struct {
    model.Params
    // 竞价请求
    bidRequest   *BidRequest
}

// 初始化YunosAdmarketAdBidRequest对象
func NewYunosAdmarketAdBidRequest() *YunosAdmarketAdBidRequest{
    return &YunosAdmarketAdBidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAdmarketAdBidRequest) GetApiMethodName() string {
    return "yunos.admarket.ad.bid"
}

// IRequest interface 方法, 获取API参数
func (r YunosAdmarketAdBidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidRequest Setter
// 竞价请求
func (r *YunosAdmarketAdBidRequest) SetBidRequest(bidRequest *BidRequest) error {
    r.bidRequest = bidRequest
    r.Set("bid_request", bidRequest)
    return nil
}

// BidRequest Getter
func (r YunosAdmarketAdBidRequest) GetBidRequest() *BidRequest {
    return r.bidRequest
}
