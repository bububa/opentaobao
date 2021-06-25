package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询店仓作业单据清单 （库存对账辅助）-回流单 APIResponse
alibaba.wdk.ums.order.get

查询店仓作业单据清单 （库存对账辅助）-回流单
*/
type AlibabaWdkUmsOrderGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsOrderGetResponse `json:"alibaba_wdk_ums_order_get_response,omitempty"`
}

type AlibabaWdkUmsOrderGetResponse struct {

    // result
    Result   *UtmsPageResult `json:"result,omitempty"`

}
