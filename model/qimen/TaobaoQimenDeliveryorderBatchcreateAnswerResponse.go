package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建结果通知接口(批量) APIResponse
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
*/
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenDeliveryorderBatchcreateAnswerResponse `json:"taobao_qimen_deliveryorder_batchcreate_answer_response,omitempty"`
}

type TaobaoQimenDeliveryorderBatchcreateAnswerResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
