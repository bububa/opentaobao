package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorProductDetailAPIResponse
获取供应链渠道中心品的详情接口（淘外分销商专用） API返回值
alibaba.ascp.channel.distributor.product.detail

此api为淘外分销的品批量查询标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorProductDetailAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorProductDetailAPIResponseModel
}

// AlibabaAscpChannelDistributorProductDetailAPIResponseModel is 获取供应链渠道中心品的详情接口（淘外分销商专用） 成功返回结果
type AlibabaAscpChannelDistributorProductDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_product_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
