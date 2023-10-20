package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse 指定分销商进行铺货(专享) API返回值
// alibaba.dchain.aoxiang.item.distribution.specify.create
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
type AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponseModel is 指定分销商进行铺货(专享) 成功返回结果
type AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_specify_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CreateItemDistributionResponse *TopResponse `json:"create_item_distribution_response,omitempty" xml:"create_item_distribution_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CreateItemDistributionResponse = nil
}

var poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse
func GetAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse() *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse {
	return poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse.Get().(*AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse 将 AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse(v *AlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionSpecifyCreateAPIResponse.Put(v)
}
