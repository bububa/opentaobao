package newretail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
setApAddressNew API返回值 
alibaba.it.ap.address.set

该接口可为ISV系统提供 ap位置信息维护的功能
*/
type AlibabaItApAddressSetAPIResponse struct {
    model.CommonResponse
    AlibabaItApAddressSetAPIResponseModel
}

// setApAddressNew 成功返回结果
type AlibabaItApAddressSetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_it_ap_address_set_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaItApAddressSetResult `json:"result,omitempty" xml:"result,omitempty"`
}
