package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
城区数据同步 APIResponse
alibaba.alihouse.newhome.region.sync

城区数据同步
*/
type AlibabaAlihouseNewhomeRegionSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeRegionSyncResponse
}

type AlibabaAlihouseNewhomeRegionSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_region_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeRegionSyncResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
