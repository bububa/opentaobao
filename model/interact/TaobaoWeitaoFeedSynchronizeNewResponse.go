package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广淘小铺isv 活动到微淘feed APIResponse
taobao.weitao.feed.synchronize.new

推广微淘互动应用活动到微淘
*/
type TaobaoWeitaoFeedSynchronizeNewAPIResponse struct {
    model.CommonResponse
    TaobaoWeitaoFeedSynchronizeNewResponse
}

type TaobaoWeitaoFeedSynchronizeNewResponse struct {
    XMLName xml.Name `xml:"weitao_feed_synchronize_new_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 增加错误信息
    
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`

    
    // 推送结果
    
    Result   *PushResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
