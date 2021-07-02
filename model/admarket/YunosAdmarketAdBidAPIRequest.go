package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosAdmarketAdBidAPIRequest 广告竞价服务 API请求
// yunos.admarket.ad.bid
//
// 广告竞价服务
type YunosAdmarketAdBidAPIRequest struct {
	model.Params
	// 竞价请求
	_bidRequest *BidRequest
}

// NewYunosAdmarketAdBidRequest 初始化YunosAdmarketAdBidAPIRequest对象
func NewYunosAdmarketAdBidRequest() *YunosAdmarketAdBidAPIRequest {
	return &YunosAdmarketAdBidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdmarketAdBidAPIRequest) GetApiMethodName() string {
	return "yunos.admarket.ad.bid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdmarketAdBidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BidRequest Setter
// 竞价请求
func (r *YunosAdmarketAdBidAPIRequest) SetBidRequest(_bidRequest *BidRequest) error {
	r._bidRequest = _bidRequest
	r.Set("bid_request", _bidRequest)
	return nil
}

// Get BidRequest Getter
func (r YunosAdmarketAdBidAPIRequest) GetBidRequest() *BidRequest {
	return r._bidRequest
}
