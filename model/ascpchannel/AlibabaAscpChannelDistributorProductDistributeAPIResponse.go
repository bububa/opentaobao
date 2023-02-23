package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductDistributeAPIResponse 分销商基于渠道产品铺货到商品 API返回值
// alibaba.ascp.channel.distributor.product.distribute
//
// 分销商基于渠道产品铺货到商品
type AlibabaAscpChannelDistributorProductDistributeAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorProductDistributeAPIResponseModel
}

// AlibabaAscpChannelDistributorProductDistributeAPIResponseModel is 分销商基于渠道产品铺货到商品 成功返回结果
type AlibabaAscpChannelDistributorProductDistributeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_product_distribute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
