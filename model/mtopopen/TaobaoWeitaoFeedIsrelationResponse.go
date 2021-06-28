package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
是否关注 APIResponse
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号
*/
type TaobaoWeitaoFeedIsrelationAPIResponse struct {
    model.CommonResponse
    TaobaoWeitaoFeedIsrelationResponse
}

type TaobaoWeitaoFeedIsrelationResponse struct {
    XMLName xml.Name `xml:"weitao_feed_isrelation_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否关注
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
