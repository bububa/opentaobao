package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓封箱回告 APIResponse
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系
*/
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_fulfill_warehouse_work_order_sealbox_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 失败返回原因
    
    RespMessage   string `json:"resp_message,omitempty" xml:"