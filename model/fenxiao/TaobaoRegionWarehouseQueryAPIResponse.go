package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionWarehouseQueryAPIResponse 查询仓库覆盖范围 API返回值
// taobao.region.warehouse.query
//
// 查询仓库覆盖范围
type TaobaoRegionWarehouseQueryAPIResponse struct {
	model.CommonResponse
	TaobaoRegionWarehouseQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionWarehouseQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionWarehouseQueryAPIResponseModel).Reset()
}

// TaobaoRegionWarehouseQueryAPIResponseModel is 查询仓库覆盖范围 成功返回结果
type TaobaoRegionWarehouseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"region_warehouse_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionWarehouseQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRegionWarehouseQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionWarehouseQueryAPIResponse)
	},
}

// GetTaobaoRegionWarehouseQueryAPIResponse 从 sync.Pool 获取 TaobaoRegionWarehouseQueryAPIResponse
func GetTaobaoRegionWarehouseQueryAPIResponse() *TaobaoRegionWarehouseQueryAPIResponse {
	return poolTaobaoRegionWarehouseQueryAPIResponse.Get().(*TaobaoRegionWarehouseQueryAPIResponse)
}

// ReleaseTaobaoRegionWarehouseQueryAPIResponse 将 TaobaoRegionWarehouseQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionWarehouseQueryAPIResponse(v *TaobaoRegionWarehouseQueryAPIResponse) {
	v.Reset()
	poolTaobaoRegionWarehouseQueryAPIResponse.Put(v)
}
