package admarket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告竞价服务 APIResponse
yunos.admarket.ad.bid

广告竞价服务
*/
type YunosAdmarketAdBidAPIResponse struct {
    model.CommonResponse
    YunosAdmarketAdBidResponse
}

type YunosAdmarketAdBidResponse struct {
    XMLName xml.Name `xml:"yunos_admarket_ad_bid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求id
    
    Id   string `json:"id,omitempty" xml:"id,omitempty"`

    
    // 返回结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 返回结果
    
    Result   *BidResponse `json:"result,omitempty" xml:"result,omitempty"`

    
    // 响应码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 是否操作成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
