package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionUpdateAPIResponse 更新商品分销内容 API返回值
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
type AlibabaDchainAoxiangItemDistributionUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemDistributionUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemDistributionUpdateAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemDistributionUpdateAPIResponseModel is 更新商品分销内容 成功返回结果
type AlibabaDchainAoxiangItemDistributionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	UpdateItemDistributionResponse *SaveItemDistributionResponse `json:"update_item_distribution_response,omitempty" xml:"update_item_distribution_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateItemDistributionResponse = nil
}

var poolAlibabaDchainAoxiangItemDistributionUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemDistributionUpdateAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemDistributionUpdateAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionUpdateAPIResponse
func GetAlibabaDchainAoxiangItemDistributionUpdateAPIResponse() *AlibabaDchainAoxiangItemDistributionUpdateAPIResponse {
	return poolAlibabaDchainAoxiangItemDistributionUpdateAPIResponse.Get().(*AlibabaDchainAoxiangItemDistributionUpdateAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemDistributionUpdateAPIResponse 将 AlibabaDchainAoxiangItemDistributionUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionUpdateAPIResponse(v *AlibabaDchainAoxiangItemDistributionUpdateAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionUpdateAPIResponse.Put(v)
}
