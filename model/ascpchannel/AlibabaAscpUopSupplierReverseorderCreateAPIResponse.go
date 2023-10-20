package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderCreateAPIResponse 商家ERP发起创建销退单服务 API返回值
// alibaba.ascp.uop.supplier.reverseorder.create
//
// 商家在收到消费者实物退货后，在ERP发起创建销退单服务
type AlibabaAscpUopSupplierReverseorderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierReverseorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel).Reset()
}

// AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel is 商家ERP发起创建销退单服务 成功返回结果
type AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_reverseorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ReverseCreateResponse *ResultWrapper `json:"reverse_create_response,omitempty" xml:"reverse_create_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReverseCreateResponse = nil
}

var poolAlibabaAscpUopSupplierReverseorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopSupplierReverseorderCreateAPIResponse)
	},
}

// GetAlibabaAscpUopSupplierReverseorderCreateAPIResponse 从 sync.Pool 获取 AlibabaAscpUopSupplierReverseorderCreateAPIResponse
func GetAlibabaAscpUopSupplierReverseorderCreateAPIResponse() *AlibabaAscpUopSupplierReverseorderCreateAPIResponse {
	return poolAlibabaAscpUopSupplierReverseorderCreateAPIResponse.Get().(*AlibabaAscpUopSupplierReverseorderCreateAPIResponse)
}

// ReleaseAlibabaAscpUopSupplierReverseorderCreateAPIResponse 将 AlibabaAscpUopSupplierReverseorderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopSupplierReverseorderCreateAPIResponse(v *AlibabaAscpUopSupplierReverseorderCreateAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopSupplierReverseorderCreateAPIResponse.Put(v)
}
