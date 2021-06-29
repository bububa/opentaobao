package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿信息同步 API返回值 
alibaba.alihouse.newhome.video.sync

接收视频信息记录
*/
type AlibabaAlihouseNewhomeVideoSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeVideoSyncResponse
}

// 视频草稿信息同步 成功返回结果
type AlibabaAlihouseNewhomeVideoSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_video_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomeVideoSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
