package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消发货 APIResponse
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态
*/
type TaobaoRdcAligeniusSendgoodsCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"rdc_aligenius_sendgoods_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoRdcAligeniusSendgoodsCancelResult `json:"result,omitempty" xml:"