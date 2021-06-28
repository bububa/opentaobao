package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
每日优鲜仓作业单回传接口 APIResponse
alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback

家乐福仓作业单回传接口
*/
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_fulfill_missfresh_warehouse_work_order_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"