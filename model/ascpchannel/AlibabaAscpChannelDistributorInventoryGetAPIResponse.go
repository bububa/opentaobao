package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorinventorygetAPIResponse 链渠道中心淘外库存查询 API返回值
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
type AlibabaascpchanneldistributorinventorygetAPIResponse struct {
	model.CommonResponse
	AlibabaascpchanneldistributorinventorygetAPIResponseModel
}

// AlibabaascpchanneldistributorinventorygetAPIResponseModel is 链渠道中心淘外库存查询 成功返回结果
type AlibabaascpchanneldistributorinventorygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_inventory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaascpchanneldistributorinventorygetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
