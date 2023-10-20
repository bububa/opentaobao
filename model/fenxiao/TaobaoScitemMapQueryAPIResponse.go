package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemMapQueryAPIResponse 查找IC商品或分销商品与后端商品的关联信息 API返回值
// taobao.scitem.map.query
//
// 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryAPIResponse struct {
	model.CommonResponse
	TaobaoScitemMapQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemMapQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemMapQueryAPIResponseModel).Reset()
}

// TaobaoScitemMapQueryAPIResponseModel is 查找IC商品或分销商品与后端商品的关联信息 成功返回结果
type TaobaoScitemMapQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_map_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后端商品映射列表
	ScItemMaps []ScItemMap `json:"sc_item_maps,omitempty" xml:"sc_item_maps>sc_item_map,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemMapQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScItemMaps = m.ScItemMaps[:0]
}

var poolTaobaoScitemMapQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemMapQueryAPIResponse)
	},
}

// GetTaobaoScitemMapQueryAPIResponse 从 sync.Pool 获取 TaobaoScitemMapQueryAPIResponse
func GetTaobaoScitemMapQueryAPIResponse() *TaobaoScitemMapQueryAPIResponse {
	return poolTaobaoScitemMapQueryAPIResponse.Get().(*TaobaoScitemMapQueryAPIResponse)
}

// ReleaseTaobaoScitemMapQueryAPIResponse 将 TaobaoScitemMapQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemMapQueryAPIResponse(v *TaobaoScitemMapQueryAPIResponse) {
	v.Reset()
	poolTaobaoScitemMapQueryAPIResponse.Put(v)
}
