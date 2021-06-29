package aliexpress

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步帖子 APIResponse
aliexpress.feed.post.publish

站外平台同步帖子至AE FEED域
*/
type AliexpressFeedPostPublishAPIResponse struct {
    model.CommonResponse
    AliexpressFeedPostPublishResponse
}

type AliexpressFeedPostPublishResponse struct {
    XMLName xml.Name `xml:"aliexpress_feed_post_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
