package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemCarturlGetAPIResponse 加购URL获取 API返回值
// taobao.item.carturl.get
//
// 获取加购URL，支持添加商品到购物车
type TaobaoItemCarturlGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemCarturlGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemCarturlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemCarturlGetAPIResponseModel).Reset()
}

// TaobaoItemCarturlGetAPIResponseModel is 加购URL获取 成功返回结果
type TaobaoItemCarturlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_carturl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加购的URL地址
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemCarturlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoItemCarturlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemCarturlGetAPIResponse)
	},
}

// GetTaobaoItemCarturlGetAPIResponse 从 sync.Pool 获取 TaobaoItemCarturlGetAPIResponse
func GetTaobaoItemCarturlGetAPIResponse() *TaobaoItemCarturlGetAPIResponse {
	return poolTaobaoItemCarturlGetAPIResponse.Get().(*TaobaoItemCarturlGetAPIResponse)
}

// ReleaseTaobaoItemCarturlGetAPIResponse 将 TaobaoItemCarturlGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemCarturlGetAPIResponse(v *TaobaoItemCarturlGetAPIResponse) {
	v.Reset()
	poolTaobaoItemCarturlGetAPIResponse.Put(v)
}
