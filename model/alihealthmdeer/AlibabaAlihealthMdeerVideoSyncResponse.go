package alihealthmdeer

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
合作伙伴视频同步给医知鹿接口 APIResponse
alibaba.alihealth.mdeer.video.sync

合伙做伴内容同步接口，用来视频同步
*/
type AlibabaAlihealthMdeerVideoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMdeerVideoSyncResponse
}

type AlibabaAlihealthMdeerVideoSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_mdeer_video_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
