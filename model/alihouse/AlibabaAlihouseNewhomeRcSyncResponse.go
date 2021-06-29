package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里房产图文草稿信息同步 APIResponse
alibaba.alihouse.newhome.rc.sync

接收图文草稿信息
*/
type AlibabaAlihouseNewhomeRcSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeRcSyncResponse
}

type AlibabaAlihouseNewhomeRcSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeRcSyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
