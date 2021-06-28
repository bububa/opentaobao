package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
是否关注 APIResponse
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号
*/
type TaobaoWeitaoFeedIsrelationAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeitaoFeedIsrelationResponse `json:"weitao_feed_isrelation_response,omitempty"` 
    TaobaoWeitaoFeedIsrelationResponse
}

/* model for simplify = false
type TaobaoWeitaoFeedIsrelationResponse struct {

    // 是否关注
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type TaobaoWeitaoFeedIsrelationResponse struct {

    // 是否关注
    Result   int64 `json:"result,omitempty"`

}
