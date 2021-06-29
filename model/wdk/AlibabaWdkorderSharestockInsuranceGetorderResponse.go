package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保消息获取 API返回值 
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceGetorderAPIResponse struct {
    model.CommonResponse
    AlibabaWdkorderSharestockInsuranceGetorderResponse
}

// 共享库存订单投保消息获取 成功返回结果
type AlibabaWdkorderSharestockInsuranceGetorderResponse struct {
    XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_getorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *MaochaoOrderInsuranceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
