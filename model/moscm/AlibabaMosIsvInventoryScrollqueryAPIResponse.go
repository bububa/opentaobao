package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosIsvInventoryScrollqueryAPIResponse 滚动查询库存数据 API返回值
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
type AlibabaMosIsvInventoryScrollqueryAPIResponse struct {
	model.CommonResponse
	AlibabaMosIsvInventoryScrollqueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosIsvInventoryScrollqueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosIsvInventoryScrollqueryAPIResponseModel).Reset()
}

// AlibabaMosIsvInventoryScrollqueryAPIResponseModel is 滚动查询库存数据 成功返回结果
type AlibabaMosIsvInventoryScrollqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_isv_inventory_scrollquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MosScrollQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosIsvInventoryScrollqueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosIsvInventoryScrollqueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosIsvInventoryScrollqueryAPIResponse)
	},
}

// GetAlibabaMosIsvInventoryScrollqueryAPIResponse 从 sync.Pool 获取 AlibabaMosIsvInventoryScrollqueryAPIResponse
func GetAlibabaMosIsvInventoryScrollqueryAPIResponse() *AlibabaMosIsvInventoryScrollqueryAPIResponse {
	return poolAlibabaMosIsvInventoryScrollqueryAPIResponse.Get().(*AlibabaMosIsvInventoryScrollqueryAPIResponse)
}

// ReleaseAlibabaMosIsvInventoryScrollqueryAPIResponse 将 AlibabaMosIsvInventoryScrollqueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosIsvInventoryScrollqueryAPIResponse(v *AlibabaMosIsvInventoryScrollqueryAPIResponse) {
	v.Reset()
	poolAlibabaMosIsvInventoryScrollqueryAPIResponse.Put(v)
}
