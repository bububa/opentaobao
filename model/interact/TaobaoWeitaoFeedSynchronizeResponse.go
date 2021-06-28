package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed APIResponse
taobao.weitao.feed.synchronize

推广淘小铺isv 活动到微淘feed
*/
type TaobaoWeitaoFeedSynchronizeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"weitao_feed_synchronize_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 同步到微淘成功与否
    
    Result   bool `json:"result,omitempty" xml:"