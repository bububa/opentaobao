package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AG退货入仓状态写接口 APIResponse
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
type TaobaoNextoneLogisticsWarehouseUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"nextone_logistics_warehouse_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // resultData
    
    ResultData   string `json:"result_data,omitempty" xml:"