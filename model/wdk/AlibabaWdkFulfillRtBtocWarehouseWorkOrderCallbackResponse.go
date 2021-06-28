package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大润发B2C仓作业单回传接口 APIResponse
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口
*/
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_fulfill_rt_btoc_warehouse_work_order_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"