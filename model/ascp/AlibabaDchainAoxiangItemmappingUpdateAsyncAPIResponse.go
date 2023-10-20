package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse 创建/更新商货品关联关系 API返回值
// alibaba.dchain.aoxiang.itemmapping.update.async
//
// 创建/更新商货品关联关系
type AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponseModel is 创建/更新商货品关联关系 成功返回结果
type AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_itemmapping_update_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	ItemMappingUpdateResponse *ItemMappingUpdateAsyncResponse `json:"item_mapping_update_response,omitempty" xml:"item_mapping_update_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemMappingUpdateResponse = nil
}

var poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse
func GetAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse() *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse {
	return poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse.Get().(*AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse 将 AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse(v *AlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemmappingUpdateAsyncAPIResponse.Put(v)
}
