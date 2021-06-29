package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息 APIResponse
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用
*/
type TmallServicecenterServicestoreUpdateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicestoreUpdateResponse
}

type TmallServicecenterServicestoreUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicestore_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 方法调用结果
    
    Result   *TmallServicecenterServicestoreUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
