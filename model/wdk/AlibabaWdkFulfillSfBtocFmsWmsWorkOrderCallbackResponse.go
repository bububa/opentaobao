package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
顺丰仓作业回传 APIResponse
alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback

顺丰仓作业单回传接口
*/
type AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_fulfill_sf_btoc_fms_wms_work_order_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"