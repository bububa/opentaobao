package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorInventoryListGetAPIResponse 批量查询渠道库存 API返回值
// alibaba.ascp.channel.distributor.inventory.list.get
//
// 淘外分销批量查询渠道产品的库存
type AlibabaAscpChannelDistributorInventoryListGetAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorInventoryListGetAPIResponseModel
}

// AlibabaAscpChannelDistributorInventoryListGetAPIResponseModel is 批量查询渠道库存 成功返回结果
type AlibabaAscpChannelDistributorInventoryListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_inventory_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpChannelDistributorInventoryListGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
