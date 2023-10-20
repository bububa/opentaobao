package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsInventoryGetinventorysAPIResponse 可售库存查询 API返回值
// alibaba.mos.goods.inventory.getinventorys
//
// 查询商品的可售、在库和占库数量
type AlibabaMosGoodsInventoryGetinventorysAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsInventoryGetinventorysAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsInventoryGetinventorysAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsInventoryGetinventorysAPIResponseModel).Reset()
}

// AlibabaMosGoodsInventoryGetinventorysAPIResponseModel is 可售库存查询 成功返回结果
type AlibabaMosGoodsInventoryGetinventorysAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_inventory_getinventorys_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据
	Datas []VirtualInventoryDto `json:"datas,omitempty" xml:"datas>virtual_inventory_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsInventoryGetinventorysAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolAlibabaMosGoodsInventoryGetinventorysAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsInventoryGetinventorysAPIResponse)
	},
}

// GetAlibabaMosGoodsInventoryGetinventorysAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsInventoryGetinventorysAPIResponse
func GetAlibabaMosGoodsInventoryGetinventorysAPIResponse() *AlibabaMosGoodsInventoryGetinventorysAPIResponse {
	return poolAlibabaMosGoodsInventoryGetinventorysAPIResponse.Get().(*AlibabaMosGoodsInventoryGetinventorysAPIResponse)
}

// ReleaseAlibabaMosGoodsInventoryGetinventorysAPIResponse 将 AlibabaMosGoodsInventoryGetinventorysAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsInventoryGetinventorysAPIResponse(v *AlibabaMosGoodsInventoryGetinventorysAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsInventoryGetinventorysAPIResponse.Put(v)
}
