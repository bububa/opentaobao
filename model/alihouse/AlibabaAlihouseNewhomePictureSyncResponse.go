package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片数据同步 API返回值 
alibaba.alihouse.newhome.picture.sync

图片数据同步
*/
type AlibabaAlihouseNewhomePictureSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomePictureSyncResponse
}

// 图片数据同步 成功返回结果
type AlibabaAlihouseNewhomePictureSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_picture_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihouseNewhomePictureSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
