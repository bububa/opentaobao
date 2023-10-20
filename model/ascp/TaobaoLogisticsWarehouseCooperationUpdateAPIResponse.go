package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationUpdateAPIResponse 合作商家信息同步 API返回值
// taobao.logistics.warehouse.cooperation.update
//
// 合作商家信息同步
type TaobaoLogisticsWarehouseCooperationUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWarehouseCooperationUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCooperationUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWarehouseCooperationUpdateAPIResponseModel).Reset()
}

// TaobaoLogisticsWarehouseCooperationUpdateAPIResponseModel is 合作商家信息同步 成功返回结果
type TaobaoLogisticsWarehouseCooperationUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_warehouse_cooperation_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 合作商家信息同步出参
	WarehouseCooperationUpdateResponse *WarehouseCooperationUpdateResponse `json:"warehouse_cooperation_update_response,omitempty" xml:"warehouse_cooperation_update_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWarehouseCooperationUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarehouseCooperationUpdateResponse = nil
}

var poolTaobaoLogisticsWarehouseCooperationUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWarehouseCooperationUpdateAPIResponse)
	},
}

// GetTaobaoLogisticsWarehouseCooperationUpdateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWarehouseCooperationUpdateAPIResponse
func GetTaobaoLogisticsWarehouseCooperationUpdateAPIResponse() *TaobaoLogisticsWarehouseCooperationUpdateAPIResponse {
	return poolTaobaoLogisticsWarehouseCooperationUpdateAPIResponse.Get().(*TaobaoLogisticsWarehouseCooperationUpdateAPIResponse)
}

// ReleaseTaobaoLogisticsWarehouseCooperationUpdateAPIResponse 将 TaobaoLogisticsWarehouseCooperationUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWarehouseCooperationUpdateAPIResponse(v *TaobaoLogisticsWarehouseCooperationUpdateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWarehouseCooperationUpdateAPIResponse.Put(v)
}
