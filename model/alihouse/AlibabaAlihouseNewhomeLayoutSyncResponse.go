package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
房通户型数据同步 APIResponse
alibaba.alihouse.newhome.layout.sync

房通户型数据同步
*/
type AlibabaAlihouseNewhomeLayoutSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeLayoutSyncResponse
}

type AlibabaAlihouseNewhomeLayoutSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_layout_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeLayoutSyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
