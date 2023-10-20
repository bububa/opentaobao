package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSalesOrderConfirmAPIResponse 渠道销售单确认收货 API返回值
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
type AlibabaAscpChannelSalesOrderConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSalesOrderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSalesOrderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSalesOrderConfirmAPIResponseModel).Reset()
}

// AlibabaAscpChannelSalesOrderConfirmAPIResponseModel is 渠道销售单确认收货 成功返回结果
type AlibabaAscpChannelSalesOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_sales_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	ConfirmOrderResponse *ResultWrapper `json:"confirm_order_response,omitempty" xml:"confirm_order_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSalesOrderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ConfirmOrderResponse = nil
}

var poolAlibabaAscpChannelSalesOrderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSalesOrderConfirmAPIResponse)
	},
}

// GetAlibabaAscpChannelSalesOrderConfirmAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSalesOrderConfirmAPIResponse
func GetAlibabaAscpChannelSalesOrderConfirmAPIResponse() *AlibabaAscpChannelSalesOrderConfirmAPIResponse {
	return poolAlibabaAscpChannelSalesOrderConfirmAPIResponse.Get().(*AlibabaAscpChannelSalesOrderConfirmAPIResponse)
}

// ReleaseAlibabaAscpChannelSalesOrderConfirmAPIResponse 将 AlibabaAscpChannelSalesOrderConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSalesOrderConfirmAPIResponse(v *AlibabaAscpChannelSalesOrderConfirmAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSalesOrderConfirmAPIResponse.Put(v)
}
