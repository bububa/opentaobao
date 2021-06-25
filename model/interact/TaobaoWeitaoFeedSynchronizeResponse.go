package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed APIResponse
taobao.weitao.feed.synchronize

推广淘小铺isv 活动到微淘feed
*/
type TaobaoWeitaoFeedSynchronizeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWeitaoFeedSynchronizeResponse `json:"taobao_weitao_feed_synchronize_response,omitempty"`
}

type TaobaoWeitaoFeedSynchronizeResponse struct {

    // 同步到微淘成功与否
    Result   bool `json:"result,omitempty"`

    // 增加错误信息
    Errmsg   string `json:"errmsg,omitempty"`

}
