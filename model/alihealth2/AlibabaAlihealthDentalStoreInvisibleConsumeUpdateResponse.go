package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店无隐形消费签约 APIResponse
alibaba.alihealth.dental.store.invisible.consume.update

门店无隐形消费签约
*/
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDentalStoreInvisibleConsumeUpdateResponse
}

type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_dental_store_invisible_consume_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
