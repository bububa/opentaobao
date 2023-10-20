package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse 供应链外部订单创建前校验接口 API返回值
// alibaba.ascp.channel.sales.order.before.check
//
// 供应链外部订单创建前校验接口
type AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSalesOrderBeforeCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSalesOrderBeforeCheckAPIResponseModel).Reset()
}

// AlibabaAscpChannelSalesOrderBeforeCheckAPIResponseModel is 供应链外部订单创建前校验接口 成功返回结果
type AlibabaAscpChannelSalesOrderBeforeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_sales_order_before_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	OrderCheckResponse *ResultWrapper `json:"order_check_response,omitempty" xml:"order_check_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSalesOrderBeforeCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderCheckResponse = nil
}

var poolAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse)
	},
}

// GetAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse
func GetAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse() *AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse {
	return poolAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse.Get().(*AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse)
}

// ReleaseAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse 将 AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse(v *AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSalesOrderBeforeCheckAPIResponse.Put(v)
}
