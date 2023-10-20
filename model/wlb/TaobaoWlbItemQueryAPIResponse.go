package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemQueryAPIResponse 分页查询商品 API返回值
// taobao.wlb.item.query
//
// 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbItemQueryAPIResponseModel).Reset()
}

// TaobaoWlbItemQueryAPIResponseModel is 分页查询商品 成功返回结果
type TaobaoWlbItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品信息列表
	ItemList []WlbItem `json:"item_list,omitempty" xml:"item_list>wlb_item,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemList = m.ItemList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbItemQueryAPIResponse)
	},
}

// GetTaobaoWlbItemQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbItemQueryAPIResponse
func GetTaobaoWlbItemQueryAPIResponse() *TaobaoWlbItemQueryAPIResponse {
	return poolTaobaoWlbItemQueryAPIResponse.Get().(*TaobaoWlbItemQueryAPIResponse)
}

// ReleaseTaobaoWlbItemQueryAPIResponse 将 TaobaoWlbItemQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbItemQueryAPIResponse(v *TaobaoWlbItemQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbItemQueryAPIResponse.Put(v)
}
