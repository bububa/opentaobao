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
func (r YunosAdmarketAdBidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAdmarketAdBidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidRequest is BidRequest Setter
// 竞价请求
func (r *YunosAdmarketAdBidAPIRequest) SetBidRequest(_bidRequest *BidRequest) error {
	r._bidRequest = _bidRequest
	r.Set("bid_request", _bidRequest)
	return nil
}

// GetBidRequest BidRequest Getter
func (r YunosAdmarketAdBidAPIRequest) GetBidRequest() *BidRequest {
	return r._bidRequest
}
