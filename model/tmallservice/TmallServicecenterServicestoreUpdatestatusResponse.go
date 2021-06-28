package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
网点/门店状态修改 APIResponse
tmall.servicecenter.servicestore.updatestatus

修改网点/门店状态
*/
type TmallServicecenterServicestoreUpdatestatusAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_servicestore_updatestatus_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 方法调用结果
    
    Result   *TmallServicecenterServicestoreUpdatestatusResult `json:"result,omitempty" xml:"