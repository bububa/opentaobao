package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsInventoryProfitlossGetAPIResponse 通过订单列表批量获取库存损益单信息 API返回值
// taobao.wlb.wms.inventory.profitloss.get
//
// 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsInventoryProfitlossGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsInventoryProfitlossGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsInventoryProfitlossGetAPIResponseModel).Reset()
}

// TaobaoWlbWmsInventoryProfitlossGetAPIResponseModel is 通过订单列表批量获取库存损益单信息 成功返回结果
type TaobaoWlbWmsInventoryProfitlossGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_inventory_profitloss_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 损益信息
	ProfitLossInfo *CainiaoInventoryProfitlossProfitlossinfo `json:"profit_loss_info,omitempty" xml:"profit_loss_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsInventoryProfitlossGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProfitLossInfo = nil
}

var poolTaobaoWlbWmsInventoryProfitlossGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsInventoryProfitlossGetAPIResponse)
	},
}

// GetTaobaoWlbWmsInventoryProfitlossGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsInventoryProfitlossGetAPIResponse
func GetTaobaoWlbWmsInventoryProfitlossGetAPIResponse() *TaobaoWlbWmsInventoryProfitlossGetAPIResponse {
	return poolTaobaoWlbWmsInventoryProfitlossGetAPIResponse.Get().(*TaobaoWlbWmsInventoryProfitlossGetAPIResponse)
}

// ReleaseTaobaoWlbWmsInventoryProfitlossGetAPIResponse 将 TaobaoWlbWmsInventoryProfitlossGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsInventoryProfitlossGetAPIResponse(v *TaobaoWlbWmsInventoryProfitlossGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsInventoryProfitlossGetAPIResponse.Put(v)
}
