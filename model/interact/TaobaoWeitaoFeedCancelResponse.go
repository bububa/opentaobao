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
    Response *TaobaoWeitaoFeedCancelResponse `json:"taobao_weitao_feed_cancel_response,omitempty"`
}

type TaobaoWeitaoFeedCancelResponse struct {

    // 推送结果
    Result   *PushResult `json:"result,omitempty"`

}
