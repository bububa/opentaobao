package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse 商家仓自营配电子面单取号 API返回值
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
type AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSelfSupplierWaybillQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopSelfSupplierWaybillQueryAPIResponseModel).Reset()
}

// AlibabaAscpUopSelfSupplierWaybillQueryAPIResponseModel is 商家仓自营配电子面单取号 成功返回结果
type AlibabaAscpUopSelfSupplierWaybillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_self_supplier_waybill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	WaybillQueryResponse *ResultWrapper `json:"waybill_query_response,omitempty" xml:"waybill_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopSelfSupplierWaybillQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WaybillQueryResponse = nil
}

var poolAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse)
	},
}

// GetAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse
func GetAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse() *AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse {
	return poolAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse.Get().(*AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse)
}

// ReleaseAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse 将 AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse(v *AlibabaAscpUopSelfSupplierWaybillQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopSelfSupplierWaybillQueryAPIResponse.Put(v)
}
