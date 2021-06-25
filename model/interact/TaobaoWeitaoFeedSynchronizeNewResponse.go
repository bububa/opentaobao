package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed APIResponse
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘
*/
type TaobaoWeitaoFeedSynchronizeNewAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWeitaoFeedSynchronizeNewResponse `json:"taobao_weitao_feed_synchronize_new_response,omitempty"`
}

type TaobaoWeitaoFeedSynchronizeNewResponse struct {

    // 增加错误信息
    Errmsg   string `json:"errmsg,omitempty"`

    // 推送结果
    Result   *PushResult `json:"result,omitempty"`

}
