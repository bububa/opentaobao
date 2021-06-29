package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核提交审核结果 APIResponse
yunos.tvpubadmin.content.single.video.submitauditresult

单视频审核提交审核结果
*/
type YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentSingleVideoSubmitauditresultResponse
}

type YunosTvpubadminContentSingleVideoSubmitauditresultResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_single_video_submitauditresult_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    Object   bool `json:"object,omitempty" xml:"object,omitempty"`

    
}
