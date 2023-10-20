package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseResourceUpdateAPIResponse 服务商新建/更新仓 API返回值
// taobao.logistics.warehouse.resource.update
//
// 服务商新建/更新仓
type TaobaoLogisticsWarehouseResourceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseResourceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseResourceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWarehouseResourceUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsWarehouseResourceUpdateAPIResponseModel is 服务商新建/更新仓 成功返回结果
type TaobaoLogisticsWarehouseResourceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_resource_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建/更新仓出参
	WarehouseUpdateResponse *WarehouseUpdateResponse `json:"warehouse_update_response,omitempty" xml:"warehouse_update_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseResourceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarehouseUpdateResponse = nil
}

var poolTaobaoLogisticsWarehouseResourceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWarehouseResourceUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsWarehouseResourceUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWarehouseResourceUpdateAPIResponse
func GetTaobaoLogisticsWarehouseResourceUpdateAPIResponse() *TaobaoLogisticsWarehouseResourceUpdateAPIResponse {
	return poolTaobaoLogisticsWarehouseResourceUpdateAPIResponse.Get().(*TaobaoLogisticsWarehouseResourceUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsWarehouseResourceUpdateAPIResponse 将 TaobaoLogisticsWarehouseResourceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWarehouseResourceUpdateAPIResponse(v *TaobaoLogisticsWarehouseResourceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWarehouseResourceUpdateAPIResponse.Put(v)
}
