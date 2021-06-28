package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消广播在timeline、广场中展示 APIResponse
taobao.weitao.feed.cancel

取消广播在timeline和广场中的展示。
*/
type TaobaoWeitaoFeedCancelAPIResponse struct {
    model.CommonResponse
    TaobaoWeitaoFeedCancelResponse
}

type TaobaoWeitaoFeedCancelResponse struct {
    XMLName xml.Name `xml:"weitao_feed_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推送结果
    
    Result   *PushResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
