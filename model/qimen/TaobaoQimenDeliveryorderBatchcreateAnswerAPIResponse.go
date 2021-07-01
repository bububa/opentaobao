package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse
发货单创建结果通知接口(批量) API返回值
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回） */
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel
}

// TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel is 发货单创建结果通知接口(批量) 成功返回结果
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchcreate_answer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderBatchcreateAnswerResponse `json:"response,omitempty" xml:"response,omitempty"`
}
