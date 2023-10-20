package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationQueryAPIResponse 仓合作关系查询 API返回值
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
type TaobaoLogisticsWarehouseCooperationQueryAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseCooperationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCooperationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWarehouseCooperationQueryAPIResponseModel).Reset()
}

// TaobaoLogisticsWarehouseCooperationQueryAPIResponseModel is 仓合作关系查询 成功返回结果
type TaobaoLogisticsWarehouseCooperationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_cooperation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓合作关系查询出参
	WarehouseCooperationQueryResponse *WarehouseCooperationQueryResponse `json:"warehouse_cooperation_query_response,omitempty" xml:"warehouse_cooperation_query_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCooperationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarehouseCooperationQueryResponse = nil
}

var poolTaobaoLogisticsWarehouseCooperationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWarehouseCooperationQueryAPIResponse)
	},
}

// GetTaobaoLogisticsWarehouseCooperationQueryAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWarehouseCooperationQueryAPIResponse
func GetTaobaoLogisticsWarehouseCooperationQueryAPIResponse() *TaobaoLogisticsWarehouseCooperationQueryAPIResponse {
	return poolTaobaoLogisticsWarehouseCooperationQueryAPIResponse.Get().(*TaobaoLogisticsWarehouseCooperationQueryAPIResponse)
}

// ReleaseTaobaoLogisticsWarehouseCooperationQueryAPIResponse 将 TaobaoLogisticsWarehouseCooperationQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWarehouseCooperationQueryAPIResponse(v *TaobaoLogisticsWarehouseCooperationQueryAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWarehouseCooperationQueryAPIResponse.Put(v)
}
