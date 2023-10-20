package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosadmarketadbidAPIRequest 广告竞价服务 API请求
// yunos.admarket.ad.bid
//
// 广告竞价服务
type YunosadmarketadbidAPIRequest struct {
	model.Params
	// 竞价请求
	_bidRequest *BidRequest
}

// NewYunosadmarketadbidRequest 初始化YunosadmarketadbidAPIRequest对象
func NewYunosadmarketadbidRequest() *YunosadmarketadbidAPIRequest {
	return &YunosadmarketadbidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosadmarketadbidAPIRequest) GetApiMethodName() string {
	return "yunos.admarket.ad.bid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosadmarketadbidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosadmarketadbidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidRequest is BidRequest Setter
// 竞价请求
func (r *YunosadmarketadbidAPIRequest) SetBidRequest(_bidRequest *BidRequest) error {
	r._bidRequest = _bidRequest
	r.Set("bid_request", _bidRequest)
	return nil
}

// GetBidRequest BidRequest Getter
func (r YunosadmarketadbidAPIRequest) GetBidRequest() *BidRequest {
	return r._bidRequest
}
