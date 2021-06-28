package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建门店 APIResponse
tmall.servicecenter.servicestore.create

用于创建门店/网点。多个业务共用
*/
type TmallServicecenterServicestoreCreateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicestoreCreateResponse
}

type TmallServicecenterServicestoreCreateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicestore_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 方法调用结果
    
    Result   *TmallServicecenterServicestoreCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
