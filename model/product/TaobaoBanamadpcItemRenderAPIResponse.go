package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemRenderAPIResponse 新发商品发布页 API返回值
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
type TaobaoBanamadpcItemRenderAPIResponse struct {
	model.CommonResponse
	TaobaoBanamadpcItemRenderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemRenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBanamadpcItemRenderAPIResponseModel).Reset()
}

// TaobaoBanamadpcItemRenderAPIResponseModel is 新发商品发布页 成功返回结果
type TaobaoBanamadpcItemRenderAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_render_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *TaobaoBanamadpcItemRenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemRenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTaobaoBanamadpcItemRenderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemRenderAPIResponse)
	},
}

// GetTaobaoBanamadpcItemRenderAPIResponse 从 sync.Pool 获取 TaobaoBanamadpcItemRenderAPIResponse
func GetTaobaoBanamadpcItemRenderAPIResponse() *TaobaoBanamadpcItemRenderAPIResponse {
	return poolTaobaoBanamadpcItemRenderAPIResponse.Get().(*TaobaoBanamadpcItemRenderAPIResponse)
}

// ReleaseTaobaoBanamadpcItemRenderAPIResponse 将 TaobaoBanamadpcItemRenderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBanamadpcItemRenderAPIResponse(v *TaobaoBanamadpcItemRenderAPIResponse) {
	v.Reset()
	poolTaobaoBanamadpcItemRenderAPIResponse.Put(v)
}
