package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeAPIResponse 库存状态同步接口 API返回值
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
type TaobaoQimenInventorySynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventorySynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenInventorySynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenInventorySynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenInventorySynchronizeAPIResponseModel is 库存状态同步接口 成功返回结果
type TaobaoQimenInventorySynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventorySynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenInventorySynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenInventorySynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventorySynchronizeAPIResponse)
	},
}

// GetTaobaoQimenInventorySynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenInventorySynchronizeAPIResponse
func GetTaobaoQimenInventorySynchronizeAPIResponse() *TaobaoQimenInventorySynchronizeAPIResponse {
	return poolTaobaoQimenInventorySynchronizeAPIResponse.Get().(*TaobaoQimenInventorySynchronizeAPIResponse)
}

// ReleaseTaobaoQimenInventorySynchronizeAPIResponse 将 TaobaoQimenInventorySynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenInventorySynchronizeAPIResponse(v *TaobaoQimenInventorySynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenInventorySynchronizeAPIResponse.Put(v)
}
