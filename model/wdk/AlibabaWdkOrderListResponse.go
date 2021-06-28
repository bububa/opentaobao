package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单拉取 APIResponse
alibaba.wdk.order.list

五道口交易订单拉取接口
*/
type AlibabaWdkOrderListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_order_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据
    
    Result   *AlibabaWdkOrderListResult `json:"result,omitempty" xml:"