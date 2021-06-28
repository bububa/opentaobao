package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询店仓作业单据清单 （库存对账辅助）-回流单 APIResponse
alibaba.wdk.ums.order.get

查询店仓作业单据清单 （库存对账辅助）-回流单
*/
type AlibabaWdkUmsOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsOrderGetResponse
}

type AlibabaWdkUmsOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_order_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
