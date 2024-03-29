package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse 逆向销退入库单到仓结果回告 API返回值
// alibaba.ascp.uop.supplier.reverseorder.instorage.result
//
// ERP回告销退入库单到仓信息回告
type AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponseModel).Reset()
}

// AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponseModel is 逆向销退入库单到仓结果回告 成功返回结果
type AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_reverseorder_instorage_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	InstorageResultResponse *ResultWrapper `json:"instorage_result_response,omitempty" xml:"instorage_result_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InstorageResultResponse = nil
}

var poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse)
	},
}

// GetAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse 从 sync.Pool 获取 AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse
func GetAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse() *AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse {
	return poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse.Get().(*AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse)
}

// ReleaseAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse 将 AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse(v *AlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopSupplierReverseorderInstorageResultAPIResponse.Put(v)
}
