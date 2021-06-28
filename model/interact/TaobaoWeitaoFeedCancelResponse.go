package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消广播在timeline、广场中展示 APIResponse
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。
*/
type TaobaoWeitaoFeedCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeitaoFeedCancelResponse `json:"weitao_feed_cancel_response,omitempty"` 
    TaobaoWeitaoFeedCancelResponse
}

/* model for simplify = false
type TaobaoWeitaoFeedCancelResponse struct {

    // 推送结果
    
    Result  *struct {
        PushResult  *PushResult `json:"push_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWeitaoFeedCancelResponse struct {

    // 推送结果
    Result   *PushResult `json:"result,omitempty"`

}
