package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松视频审核记录查询 APIResponse
yunos.tvpubadmin.content.video.getauditlist

迎客松视频审核记录查询
*/
type YunosTvpubadminContentVideoGetauditlistAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentVideoGetauditlistResponse
}

type YunosTvpubadminContentVideoGetauditlistResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_video_getauditlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 视频审核列表，json格式
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
