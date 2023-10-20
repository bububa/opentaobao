package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemGetAPIResponse 根据id查询商品 API返回值
// taobao.scitem.get
//
// 根据id查询商品
type TaobaoScitemGetAPIResponse struct {
	model.CommonResponse
	TaobaoScitemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemGetAPIResponseModel).Reset()
}

// TaobaoScitemGetAPIResponseModel is 根据id查询商品 成功返回结果
type TaobaoScitemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后端商品
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScItem = nil
}

var poolTaobaoScitemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemGetAPIResponse)
	},
}

// GetTaobaoScitemGetAPIResponse 从 sync.Pool 获取 TaobaoScitemGetAPIResponse
func GetTaobaoScitemGetAPIResponse() *TaobaoScitemGetAPIResponse {
	return poolTaobaoScitemGetAPIResponse.Get().(*TaobaoScitemGetAPIResponse)
}

// ReleaseTaobaoScitemGetAPIResponse 将 TaobaoScitemGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemGetAPIResponse(v *TaobaoScitemGetAPIResponse) {
	v.Reset()
	poolTaobaoScitemGetAPIResponse.Put(v)
}
