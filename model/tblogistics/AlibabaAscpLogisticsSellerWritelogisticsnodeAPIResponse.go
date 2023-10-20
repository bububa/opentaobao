package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse 商家配送写入物流节点 API返回值
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
type AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponseModel is 商家配送写入物流节点 成功返回结果
type AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_seller_writelogisticsnode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BatchWriteLogisticsNodeTopResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse)
	},
}

// GetAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse
func GetAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse() *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse {
	return poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse.Get().(*AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse)
}

// ReleaseAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse 将 AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse(v *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse.Put(v)
}
