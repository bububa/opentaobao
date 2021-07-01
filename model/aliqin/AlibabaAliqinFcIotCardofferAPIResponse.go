package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联网卡上订购的offer API返回值 
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer
*/
type AlibabaAliqinFcIotCardofferAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotCardofferAPIResponseModel
}

// 查询物联网卡上订购的offer 成功返回结果
type AlibabaAliqinFcIotCardofferAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardoffer_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *AlibabaAliqinFcIotCardofferResult `json:"result,omitempty" xml:"result,omitempty"`
}
