package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductQuantityUpdateAPIResponse 渠道无仓库存更新接口 API返回值
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
type TmallSupplychainChannelProductQuantityUpdateAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductQuantityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductQuantityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductQuantityUpdateAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductQuantityUpdateAPIResponseModel is 渠道无仓库存更新接口 成功返回结果
type TmallSupplychainChannelProductQuantityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_quantity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductQuantityUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductQuantityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductQuantityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductQuantityUpdateAPIResponse)
	},
}

// GetTmallSupplychainChannelProductQuantityUpdateAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductQuantityUpdateAPIResponse
func GetTmallSupplychainChannelProductQuantityUpdateAPIResponse() *TmallSupplychainChannelProductQuantityUpdateAPIResponse {
	return poolTmallSupplychainChannelProductQuantityUpdateAPIResponse.Get().(*TmallSupplychainChannelProductQuantityUpdateAPIResponse)
}

// ReleaseTmallSupplychainChannelProductQuantityUpdateAPIResponse 将 TmallSupplychainChannelProductQuantityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductQuantityUpdateAPIResponse(v *TmallSupplychainChannelProductQuantityUpdateAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductQuantityUpdateAPIResponse.Put(v)
}
