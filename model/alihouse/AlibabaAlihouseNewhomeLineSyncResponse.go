package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
环线数据同步 APIResponse
alibaba.alihouse.newhome.line.sync

环线数据同步
*/
type AlibabaAlihouseNewhomeLineSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeLineSyncResponse
}

type AlibabaAlihouseNewhomeLineSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_line_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeLineSyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
