package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemMapGetAPIResponse 根据物流宝商品ID查询商品映射关系 API返回值
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
type TaobaoWlbItemMapGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemMapGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbItemMapGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbItemMapGetAPIResponseModel).Reset()
}

// TaobaoWlbItemMapGetAPIResponseModel is 根据物流宝商品ID查询商品映射关系 成功返回结果
type TaobaoWlbItemMapGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_map_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外部商品实体
	OutEntityItemList []OutEntityItem `json:"out_entity_item_list,omitempty" xml:"out_entity_item_list>out_entity_item,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbItemMapGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OutEntityItemList = m.OutEntityItemList[:0]
	m.IsSuccess = false
}

var poolTaobaoWlbItemMapGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbItemMapGetAPIResponse)
	},
}

// GetTaobaoWlbItemMapGetAPIResponse 从 sync.Pool 获取 TaobaoWlbItemMapGetAPIResponse
func GetTaobaoWlbItemMapGetAPIResponse() *TaobaoWlbItemMapGetAPIResponse {
	return poolTaobaoWlbItemMapGetAPIResponse.Get().(*TaobaoWlbItemMapGetAPIResponse)
}

// ReleaseTaobaoWlbItemMapGetAPIResponse 将 TaobaoWlbItemMapGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbItemMapGetAPIResponse(v *TaobaoWlbItemMapGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbItemMapGetAPIResponse.Put(v)
}
