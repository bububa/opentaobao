package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿状态更新 APIResponse
alibaba.alihouse.newhome.video.changestatus

视频草稿状态更新
*/
type AlibabaAlihouseNewhomeVideoChangestatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeVideoChangestatusResponse
}

type AlibabaAlihouseNewhomeVideoChangestatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_video_changestatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeVideoChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
