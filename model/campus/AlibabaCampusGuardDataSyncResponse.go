package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卡巴数据同步 APIResponse
alibaba.campus.guard.data.sync

数据同步门禁系统
*/
type AlibabaCampusGuardDataSyncAPIResponse struct {
    model.CommonResponse
    AlibabaCampusGuardDataSyncResponse
}

type AlibabaCampusGuardDataSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_guard_data_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
