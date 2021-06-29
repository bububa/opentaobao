package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单创建接口 API返回值 
alibaba.servicecenter.workcard.create

创建服务平台工单
*/
type AlibabaServicecenterWorkcardCreateAPIResponse struct {
    model.CommonResponse
    AlibabaServicecenterWorkcardCreateResponse
}

// 服务平台工单创建接口 成功返回结果
type AlibabaServicecenterWorkcardCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_servicecenter_workcard_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回参数
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
