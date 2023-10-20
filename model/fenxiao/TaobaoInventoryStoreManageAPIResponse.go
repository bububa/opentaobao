package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryStoreManageAPIResponse 创建或更新仓库 API返回值
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
type TaobaoInventoryStoreManageAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryStoreManageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryStoreManageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryStoreManageAPIResponseModel).Reset()
}

// TaobaoInventoryStoreManageAPIResponseModel is 创建或更新仓库 成功返回结果
type TaobaoInventoryStoreManageAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_store_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	StoreList []Store `json:"store_list,omitempty" xml:"store_list>store,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryStoreManageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreList = m.StoreList[:0]
}

var poolTaobaoInventoryStoreManageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryStoreManageAPIResponse)
	},
}

// GetTaobaoInventoryStoreManageAPIResponse 从 sync.Pool 获取 TaobaoInventoryStoreManageAPIResponse
func GetTaobaoInventoryStoreManageAPIResponse() *TaobaoInventoryStoreManageAPIResponse {
	return poolTaobaoInventoryStoreManageAPIResponse.Get().(*TaobaoInventoryStoreManageAPIResponse)
}

// ReleaseTaobaoInventoryStoreManageAPIResponse 将 TaobaoInventoryStoreManageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryStoreManageAPIResponse(v *TaobaoInventoryStoreManageAPIResponse) {
	v.Reset()
	poolTaobaoInventoryStoreManageAPIResponse.Put(v)
}
