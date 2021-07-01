package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovGetLatestbidAPIRequest
获取司法拍卖最新出价数据 API请求
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据 */
type TaobaoAuctionGovGetLatestbidAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 死否包含下属法院
	_containChild bool
	// 获取最新出价条数
	_maxCount int64
}

// New
