package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsItemCombinationGetAPIResponse 查询组合商品的组合关系 API返回值
// taobao.wlb.wms.item.combination.get
//
// 查询组合商品的组合关系
type TaobaoWlbWmsItemCombinationGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsItemCombinationGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsItemCombinationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsItemCombinationGetAPIResponseModel).Reset()
}

// TaobaoWlbWmsItemCombinationGetAPIResponseModel is 查询组合商品的组合关系 成功返回结果
type TaobaoWlbWmsItemCombinationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_item_combination_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *TaobaoWlbWmsItemCombinationGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsItemCombinationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbWmsItemCombinationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsItemCombinationGetAPIResponse)
	},
}

// GetTaobaoWlbWmsItemCombinationGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsItemCombinationGetAPIResponse
func GetTaobaoWlbWmsItemCombinationGetAPIResponse() *TaobaoWlbWmsItemCombinationGetAPIResponse {
	return poolTaobaoWlbWmsItemCombinationGetAPIResponse.Get().(*TaobaoWlbWmsItemCombinationGetAPIResponse)
}

// ReleaseTaobaoWlbWmsItemCombinationGetAPIResponse 将 TaobaoWlbWmsItemCombinationGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsItemCombinationGetAPIResponse(v *TaobaoWlbWmsItemCombinationGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsItemCombinationGetAPIResponse.Put(v)
}
