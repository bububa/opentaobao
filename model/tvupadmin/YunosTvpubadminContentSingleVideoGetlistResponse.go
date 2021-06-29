package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核获取视频列表 APIResponse
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口
*/
type YunosTvpubadminContentSingleVideoGetlistAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentSingleVideoGetlistResponse
}

type YunosTvpubadminContentSingleVideoGetlistResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_single_video_getlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
