package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdmarketAdBidAPIRequest
广告竞价服务 API请求
yunos.admarket.ad.bid

广告竞价服务 */
type YunosAdmarketAdBidAPIRequest struct {
	model.Params
	// 竞价请求
	_bidRequest *BidRequest
}

// New
