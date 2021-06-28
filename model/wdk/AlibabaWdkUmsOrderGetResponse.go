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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsPageResult `json:"result,omitempty" xml:"