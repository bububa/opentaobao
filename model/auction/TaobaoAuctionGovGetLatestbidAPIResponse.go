package auction

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovGetLatestbidAPIResponse
获取司法拍卖最新出价数据 API返回值
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据 */
type TaobaoAuctionGovGetLatestbidAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionGovGetLatestbidAPIResponseModel
}

// TaobaoAuctionGovGetLatestbidAPIResponseModel is 获取司法拍卖最新出价数据 成功返回结果
type TaobaoAuctionGovGetLatestbidAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_get_latestbid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *Result4Top `json:"result,omitempty" xml:"result,omitempty"`
}
