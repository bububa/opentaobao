package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松提交视频审核结果 APIResponse
yunos.tvpubadmin.content.video.submitauditresult

迎客松提交视频审核结果
*/
type YunosTvpubadminContentVideoSubmitauditresultAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentVideoSubmitauditresultResponse
}

type YunosTvpubadminContentVideoSubmitauditresultResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_video_submitauditresult_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
