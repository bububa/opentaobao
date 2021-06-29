package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
C2M售后状态同步 API返回值 
alibaba.ihome.ctom.postsale.status.sync

供给三维家同步定制、成品商品售后进度状态
*/
type AlibabaIhomeCtomPostsaleStatusSyncAPIResponse struct {
    model.CommonResponse
    AlibabaIhomeCtomPostsaleStatusSyncResponse
}

// C2M售后状态同步 成功返回结果
type AlibabaIhomeCtomPostsaleStatusSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_ihome_ctom_postsale_status_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
