package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorPriceGetAPIResponse 链渠道中心淘外分销价格查询(分销商专用) API返回值
// alibaba.ascp.channel.distributor.price.get
//
// 此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorPriceGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorPriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorPriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelDistributorPriceGetAPIResponseModel).Reset()
}

// AlibabaAscpChannelDistributorPriceGetAPIResponseModel is 链渠道中心淘外分销价格查询(分销商专用) 成功返回结果
type AlibabaAscpChannelDistributorPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpChannelDistributorPriceGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorPriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelDistributorPriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorPriceGetAPIResponse)
	},
}

// GetAlibabaAscpChannelDistributorPriceGetAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelDistributorPriceGetAPIResponse
func GetAlibabaAscpChannelDistributorPriceGetAPIResponse() *AlibabaAscpChannelDistributorPriceGetAPIResponse {
	return poolAlibabaAscpChannelDistributorPriceGetAPIResponse.Get().(*AlibabaAscpChannelDistributorPriceGetAPIResponse)
}

// ReleaseAlibabaAscpChannelDistributorPriceGetAPIResponse 将 AlibabaAscpChannelDistributorPriceGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelDistributorPriceGetAPIResponse(v *AlibabaAscpChannelDistributorPriceGetAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelDistributorPriceGetAPIResponse.Put(v)
}
