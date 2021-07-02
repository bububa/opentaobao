package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAPIResponse 发货单创建批量接口 API返回值
// taobao.qimen.deliveryorder.batchcreate
//
// ERP调用接口，将发货信息批量推送给WMS
type TaobaoQimenDeliveryorderBatchcreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchcreateAPIResponseModel
}

// TaobaoQimenDeliveryorderBatchcreateAPIResponseModel is 发货单创建批量接口 成功返回结果
type TaobaoQimenDeliveryorderBatchcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchcreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *DeliveryOrderBatchCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
