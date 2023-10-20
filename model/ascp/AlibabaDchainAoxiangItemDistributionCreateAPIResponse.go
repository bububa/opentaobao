package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionCreateAPIResponse 选择店铺商品并进行铺货（通用） API返回值
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
type AlibabaDchainAoxiangItemDistributionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemDistributionCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemDistributionCreateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemDistributionCreateAPIResponseModel is 选择店铺商品并进行铺货（通用） 成功返回结果
type AlibabaDchainAoxiangItemDistributionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CreateItemDistributionResponse *TopResponse `json:"create_item_distribution_response,omitempty" xml:"create_item_distribution_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CreateItemDistributionResponse = nil
}

var poolAlibabaDchainAoxiangItemDistributionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemDistributionCreateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemDistributionCreateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionCreateAPIResponse
func GetAlibabaDchainAoxiangItemDistributionCreateAPIResponse() *AlibabaDchainAoxiangItemDistributionCreateAPIResponse {
	return poolAlibabaDchainAoxiangItemDistributionCreateAPIResponse.Get().(*AlibabaDchainAoxiangItemDistributionCreateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemDistributionCreateAPIResponse 将 AlibabaDchainAoxiangItemDistributionCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionCreateAPIResponse(v *AlibabaDchainAoxiangItemDistributionCreateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionCreateAPIResponse.Put(v)
}
