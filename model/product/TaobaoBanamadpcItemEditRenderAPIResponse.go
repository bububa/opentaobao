package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemEditRenderAPIResponse 编辑商品发布页 API返回值
// taobao.banamadpc.item.edit.render
//
// 巴拿马供应商通过此接口获取编辑商品发布页
type TaobaoBanamadpcItemEditRenderAPIResponse struct {
	model.CommonResponse
	TaobaoBanamadpcItemEditRenderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemEditRenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBanamadpcItemEditRenderAPIResponseModel).Reset()
}

// TaobaoBanamadpcItemEditRenderAPIResponseModel is 编辑商品发布页 成功返回结果
type TaobaoBanamadpcItemEditRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_edit_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaoBanamadpcItemEditRenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemEditRenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTaobaoBanamadpcItemEditRenderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemEditRenderAPIResponse)
	},
}

// GetTaobaoBanamadpcItemEditRenderAPIResponse 从 sync.Pool 获取 TaobaoBanamadpcItemEditRenderAPIResponse
func GetTaobaoBanamadpcItemEditRenderAPIResponse() *TaobaoBanamadpcItemEditRenderAPIResponse {
	return poolTaobaoBanamadpcItemEditRenderAPIResponse.Get().(*TaobaoBanamadpcItemEditRenderAPIResponse)
}

// ReleaseTaobaoBanamadpcItemEditRenderAPIResponse 将 TaobaoBanamadpcItemEditRenderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBanamadpcItemEditRenderAPIResponse(v *TaobaoBanamadpcItemEditRenderAPIResponse) {
	v.Reset()
	poolTaobaoBanamadpcItemEditRenderAPIResponse.Put(v)
}
