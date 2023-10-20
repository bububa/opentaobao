package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse 商家查询负卖库存 API返回值
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.query
//
// 商家根据当前接口查询负卖货品的库存
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel).Reset()
}

// AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel is 商家查询负卖库存 成功返回结果
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_negative_sale_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	AicinventoryQueryResponse *ResultWrapper `json:"aicinventory_query_response,omitempty" xml:"aicinventory_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AicinventoryQueryResponse = nil
}

var poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse)
	},
}

// GetAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse
func GetAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse() *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse {
	return poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse.Get().(*AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse 将 AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse(v *AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIResponse.Put(v)
}
