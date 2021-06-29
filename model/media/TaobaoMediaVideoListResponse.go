package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家视频列表 APIResponse
taobao.media.video.list

用于获取授权商家的视频列表
*/
type TaobaoMediaVideoListAPIResponse struct {
    model.CommonResponse
    TaobaoMediaVideoListResponse
}

type TaobaoMediaVideoListResponse struct {
    XMLName xml.Name `xml:"media_video_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *SearchResultDO `json:"result,omitempty" xml:"result,omitempty"`

    
}
