package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorInventoryGetAPIResponse 链渠道中心淘外库存查询 API返回值
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
type AlibabaAscpChannelDistributorInventoryGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorInventoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorInventoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelDistributorInventoryGetAPIResponseModel).Reset()
}

// AlibabaAscpChannelDistributorInventoryGetAPIResponseModel is 链渠道中心淘外库存查询 成功返回结果
type AlibabaAscpChannelDistributorInventoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_inventory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpChannelDistributorInventoryGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorInventoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelDistributorInventoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorInventoryGetAPIResponse)
	},
}

// GetAlibabaAscpChannelDistributorInventoryGetAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelDistributorInventoryGetAPIResponse
func GetAlibabaAscpChannelDistributorInventoryGetAPIResponse() *AlibabaAscpChannelDistributorInventoryGetAPIResponse {
	return poolAlibabaAscpChannelDistributorInventoryGetAPIResponse.Get().(*AlibabaAscpChannelDistributorInventoryGetAPIResponse)
}

// ReleaseAlibabaAscpChannelDistributorInventoryGetAPIResponse 将 AlibabaAscpChannelDistributorInventoryGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelDistributorInventoryGetAPIResponse(v *AlibabaAscpChannelDistributorInventoryGetAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelDistributorInventoryGetAPIResponse.Put(v)
}
