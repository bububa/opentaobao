package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemAddAPIResponse 发布后端商品 API返回值
// taobao.scitem.add
//
// 发布后端商品
type TaobaoScitemAddAPIResponse struct {
	model.CommonResponse
	TaobaoScitemAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemAddAPIResponseModel).Reset()
}

// TaobaoScitemAddAPIResponseModel is 发布后端商品 成功返回结果
type TaobaoScitemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后台商品信息
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScItem = nil
}

var poolTaobaoScitemAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemAddAPIResponse)
	},
}

// GetTaobaoScitemAddAPIResponse 从 sync.Pool 获取 TaobaoScitemAddAPIResponse
func GetTaobaoScitemAddAPIResponse() *TaobaoScitemAddAPIResponse {
	return poolTaobaoScitemAddAPIResponse.Get().(*TaobaoScitemAddAPIResponse)
}

// ReleaseTaobaoScitemAddAPIResponse 将 TaobaoScitemAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemAddAPIResponse(v *TaobaoScitemAddAPIResponse) {
	v.Reset()
	poolTaobaoScitemAddAPIResponse.Put(v)
}
