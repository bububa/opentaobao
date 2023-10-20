package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemCombinationGetAPIResponse 根据商品id查询商品组合关系 API返回值
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
type TaobaoWlbItemCombinationGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemCombinationGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbItemCombinationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbItemCombinationGetAPIResponseModel).Reset()
}

// TaobaoWlbItemCombinationGetAPIResponseModel is 根据商品id查询商品组合关系 成功返回结果
type TaobaoWlbItemCombinationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_combination_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 组合子商品id列表
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbItemCombinationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemIdList = m.ItemIdList[:0]
}

var poolTaobaoWlbItemCombinationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbItemCombinationGetAPIResponse)
	},
}

// GetTaobaoWlbItemCombinationGetAPIResponse 从 sync.Pool 获取 TaobaoWlbItemCombinationGetAPIResponse
func GetTaobaoWlbItemCombinationGetAPIResponse() *TaobaoWlbItemCombinationGetAPIResponse {
	return poolTaobaoWlbItemCombinationGetAPIResponse.Get().(*TaobaoWlbItemCombinationGetAPIResponse)
}

// ReleaseTaobaoWlbItemCombinationGetAPIResponse 将 TaobaoWlbItemCombinationGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbItemCombinationGetAPIResponse(v *TaobaoWlbItemCombinationGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbItemCombinationGetAPIResponse.Put(v)
}
