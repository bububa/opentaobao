package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建批量接口 APIResponse
taobao.qimen.deliveryorder.batchcreate

ERP调用接口，将发货信息批量推送给WMS
*/
type TaobaoQimenDeliveryorderBatchcreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_deliveryorder_batchcreate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *DeliveryOrderBatchCreateResponse `json:"response,omitempty" xml:"