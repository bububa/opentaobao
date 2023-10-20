package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryStoreQueryAPIResponse 查询仓库信息 API返回值
// taobao.inventory.store.query
//
// 查询商家仓信息
type TaobaoInventoryStoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryStoreQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryStoreQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryStoreQueryAPIResponseModel).Reset()
}

// TaobaoInventoryStoreQueryAPIResponseModel is 查询仓库信息 成功返回结果
type TaobaoInventoryStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓库列表
	StoreList []Store `json:"store_list,omitempty" xml:"store_list>store,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryStoreQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreList = m.StoreList[:0]
}

var poolTaobaoInventoryStoreQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryStoreQueryAPIResponse)
	},
}

// GetTaobaoInventoryStoreQueryAPIResponse 从 sync.Pool 获取 TaobaoInventoryStoreQueryAPIResponse
func GetTaobaoInventoryStoreQueryAPIResponse() *TaobaoInventoryStoreQueryAPIResponse {
	return poolTaobaoInventoryStoreQueryAPIResponse.Get().(*TaobaoInventoryStoreQueryAPIResponse)
}

// ReleaseTaobaoInventoryStoreQueryAPIResponse 将 TaobaoInventoryStoreQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryStoreQueryAPIResponse(v *TaobaoInventoryStoreQueryAPIResponse) {
	v.Reset()
	poolTaobaoInventoryStoreQueryAPIResponse.Put(v)
}
