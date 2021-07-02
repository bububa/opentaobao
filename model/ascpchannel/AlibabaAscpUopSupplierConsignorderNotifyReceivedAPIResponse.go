package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponse 商家仓物流发货推单接单回告 API返回值
// alibaba.ascp.uop.supplier.consignorder.notify.received
//
// ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
type AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponseModel
}

// AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponseModel is 商家仓物流发货推单接单回告 成功返回结果
type AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_notify_received_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ConsignorderNotifyReceivedResponse *ResultWrapper `json:"consignorder_notify_received_response,omitempty" xml:"consignorder_notify_received_response,omitempty"`
}
