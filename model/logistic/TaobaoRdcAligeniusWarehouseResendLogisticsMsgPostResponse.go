package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
补发单物流信息回传 API返回值 
taobao.rdc.aligenius.warehouse.resend.logistics.msg.post

补发单erp物流信息回传平台
*/
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse struct {
    model.CommonResponse
    TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse
}

// 补发单物流信息回传 成功返回结果
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResponse struct {
    XMLName xml.Name `xml:"rdc_aligenius_warehouse_resend_logistics_msg_post_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult `json:"result,omitempty" xml:"result,omitempty"`
}
