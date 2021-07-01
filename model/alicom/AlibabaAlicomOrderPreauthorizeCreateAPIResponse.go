package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
业务办理结果 API返回值 
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知
*/
type AlibabaAlicomOrderPreauthorizeCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel
}

// 业务办理结果 成功返回结果
type AlibabaAlicomOrderPreauthorizeCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
