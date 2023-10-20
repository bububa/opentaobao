package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse 商家仓wms取消发货反馈回告服务 API返回值
// alibaba.ascp.uop.supplier.consignorder.cancel.feedback
//
// 履约单纬度通知商家仓wms取消发货结果反馈回告服务
type AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponseModel).Reset()
}

// AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponseModel is 商家仓wms取消发货反馈回告服务 成功返回结果
type AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_cancel_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ConsignorderCancelFeedbackResponse *ResultWrapper `json:"consignorder_cancel_feedback_response,omitempty" xml:"consignorder_cancel_feedback_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ConsignorderCancelFeedbackResponse = nil
}

var poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse)
	},
}

// GetAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse
func GetAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse() *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse {
	return poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse.Get().(*AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse)
}

// ReleaseAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse 将 AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse(v *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse.Put(v)
}
