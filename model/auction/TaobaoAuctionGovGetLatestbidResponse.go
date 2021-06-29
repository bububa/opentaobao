package auction

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取司法拍卖最新出价数据 APIResponse
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据
*/
type TaobaoAuctionGovGetLatestbidAPIResponse struct {
    model.CommonResponse
    TaobaoAuctionGovGetLatestbidResponse
}

type TaobaoAuctionGovGetLatestbidResponse struct {
    XMLName xml.Name `xml:"auction_gov_get_latestbid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *Result4Top `json:"result,omitempty" xml:"result,omitempty"`

    
}
