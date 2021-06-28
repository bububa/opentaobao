package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户订单履约数据获取 APIResponse
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
type AlibabaWdkorderSharestockFulfillGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdkorder_sharestock_fulfill_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *MaochaoOrderFulfillQueryResult `json:"result,omitempty" xml:"