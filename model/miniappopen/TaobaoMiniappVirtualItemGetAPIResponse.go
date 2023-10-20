package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappVirtualItemGetAPIResponse 小程序关联虚拟商品查询 API返回值
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
type TaobaoMiniappVirtualItemGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappVirtualItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappVirtualItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappVirtualItemGetAPIResponseModel).Reset()
}

// TaobaoMiniappVirtualItemGetAPIResponseModel is 小程序关联虚拟商品查询 成功返回结果
type TaobaoMiniappVirtualItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_virtual_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	Model []MiniappItemDto `json:"model,omitempty" xml:"model>miniapp_item_dto,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 业务错误信息
	ECode int64 `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappVirtualItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
	m.ErrMessage = ""
	m.ECode = 0
	m.Suc = false
}

var poolTaobaoMiniappVirtualItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappVirtualItemGetAPIResponse)
	},
}

// GetTaobaoMiniappVirtualItemGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappVirtualItemGetAPIResponse
func GetTaobaoMiniappVirtualItemGetAPIResponse() *TaobaoMiniappVirtualItemGetAPIResponse {
	return poolTaobaoMiniappVirtualItemGetAPIResponse.Get().(*TaobaoMiniappVirtualItemGetAPIResponse)
}

// ReleaseTaobaoMiniappVirtualItemGetAPIResponse 将 TaobaoMiniappVirtualItemGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappVirtualItemGetAPIResponse(v *TaobaoMiniappVirtualItemGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappVirtualItemGetAPIResponse.Put(v)
}
