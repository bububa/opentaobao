package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务办理结果 APIResponse
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知
*/
type AlibabaAlicomOrderPreauthorizeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomOrderPreauthorizeCreateResponse
}

type AlibabaAlicomOrderPreauthorizeCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
