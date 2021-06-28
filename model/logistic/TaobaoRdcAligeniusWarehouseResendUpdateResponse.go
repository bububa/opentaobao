package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单状态回传 APIResponse
taobao.rdc.aligenius.warehouse.resend.update

补发单状态回传接口
*/
type TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusWarehouseResendUpdateResponse
}

type TaobaoRdcAligeniusWarehouseResendUpdateResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_warehouse_resend_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusWarehouseResendUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
