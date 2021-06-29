package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
内部迎客松通用服务 APIResponse
yunos.pubadmin.common.operation

内部迎客松通用服务
*/
type YunosPubadminCommonOperationAPIResponse struct {
    model.CommonResponse
    YunosPubadminCommonOperationResponse
}

type YunosPubadminCommonOperationResponse struct {
    XMLName xml.Name `xml:"yunos_pubadmin_common_operation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
