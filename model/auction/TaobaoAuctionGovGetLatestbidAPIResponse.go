package auction

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovgetlatestbidAPIResponse 获取司法拍卖最新出价数据 API返回值
// taobao.auction.gov.get.latestbid
//
// 获取司法拍卖最新出价数据
type TaobaoauctiongovgetlatestbidAPIResponse struct {
	model.CommonResponse
	TaobaoauctiongovgetlatestbidAPIResponseModel
}

// TaobaoauctiongovgetlatestbidAPIResponseModel is 获取司法拍卖最新出价数据 成功返回结果
type TaobaoauctiongovgetlatestbidAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_get_latestbid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *Result4top `json:"result,omitempty" xml:"result,omitempty"`
}
