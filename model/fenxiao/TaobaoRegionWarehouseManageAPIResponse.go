package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionWarehouseManageAPIResponse 编辑仓库覆盖范围 API返回值
// taobao.region.warehouse.manage
//
// 编辑仓库覆盖范围
type TaobaoRegionWarehouseManageAPIResponse struct {
	model.CommonResponse
	TaobaoRegionWarehouseManageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionWarehouseManageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionWarehouseManageAPIResponseModel).Reset()
}

// TaobaoRegionWarehouseManageAPIResponseModel is 编辑仓库覆盖范围 成功返回结果
type TaobaoRegionWarehouseManageAPIResponseModel struct {
	XMLName xml.Name `xml:"region_warehouse_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionWarehouseManageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRegionWarehouseManageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionWarehouseManageAPIResponse)
	},
}

// GetTaobaoRegionWarehouseManageAPIResponse 从 sync.Pool 获取 TaobaoRegionWarehouseManageAPIResponse
func GetTaobaoRegionWarehouseManageAPIResponse() *TaobaoRegionWarehouseManageAPIResponse {
	return poolTaobaoRegionWarehouseManageAPIResponse.Get().(*TaobaoRegionWarehouseManageAPIResponse)
}

// ReleaseTaobaoRegionWarehouseManageAPIResponse 将 TaobaoRegionWarehouseManageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionWarehouseManageAPIResponse(v *TaobaoRegionWarehouseManageAPIResponse) {
	v.Reset()
	poolTaobaoRegionWarehouseManageAPIResponse.Put(v)
}
