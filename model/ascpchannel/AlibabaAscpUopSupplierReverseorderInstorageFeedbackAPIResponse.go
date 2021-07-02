package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIResponse 逆向销退入库单入库结果回告 API返回值
// alibaba.ascp.uop.supplier.reverseorder.instorage.feedback
//
// ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIResponseModel
}

// AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIResponseModel is 逆向销退入库单入库结果回告 成功返回结果
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_reverseorder_instorage_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	InstorageFeedbackResponse *ResultWrapper `json:"instorage_feedback_response,omitempty" xml:"instorage_feedback_response,omitempty"`
}
