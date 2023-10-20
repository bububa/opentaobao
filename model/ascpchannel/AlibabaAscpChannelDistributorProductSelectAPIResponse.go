package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelDistributorProductSelectAPIResponse 供应链渠道中心品的选品接口（淘外分销商专用） API返回值
// alibaba.ascp.channel.distributor.product.select
//
// 此api为淘外分销的品的选品标准api，淘外分销商专用
type AlibabaAscpChannelDistributorProductSelectAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelDistributorProductSelectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorProductSelectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelDistributorProductSelectAPIResponseModel).Reset()
}

// AlibabaAscpChannelDistributorProductSelectAPIResponseModel is 供应链渠道中心品的选品接口（淘外分销商专用） 成功返回结果
type AlibabaAscpChannelDistributorProductSelectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_product_select_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 系统自动生成
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 返回是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelDistributorProductSelectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessage = ""
	m.Error = ""
	m.ResultSuccess = false
}

var poolAlibabaAscpChannelDistributorProductSelectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorProductSelectAPIResponse)
	},
}

// GetAlibabaAscpChannelDistributorProductSelectAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelDistributorProductSelectAPIResponse
func GetAlibabaAscpChannelDistributorProductSelectAPIResponse() *AlibabaAscpChannelDistributorProductSelectAPIResponse {
	return poolAlibabaAscpChannelDistributorProductSelectAPIResponse.Get().(*AlibabaAscpChannelDistributorProductSelectAPIResponse)
}

// ReleaseAlibabaAscpChannelDistributorProductSelectAPIResponse 将 AlibabaAscpChannelDistributorProductSelectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelDistributorProductSelectAPIResponse(v *AlibabaAscpChannelDistributorProductSelectAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelDistributorProductSelectAPIResponse.Put(v)
}
