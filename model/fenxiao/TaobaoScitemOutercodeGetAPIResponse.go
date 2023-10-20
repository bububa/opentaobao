package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemOutercodeGetAPIResponse 根据outerCode查询商品 API返回值
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
type TaobaoScitemOutercodeGetAPIResponse struct {
	model.CommonResponse
	TaobaoScitemOutercodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoScitemOutercodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoScitemOutercodeGetAPIResponseModel).Reset()
}

// TaobaoScitemOutercodeGetAPIResponseModel is 根据outerCode查询商品 成功返回结果
type TaobaoScitemOutercodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_outercode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后台商品
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoScitemOutercodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ScItem = nil
}

var poolTaobaoScitemOutercodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoScitemOutercodeGetAPIResponse)
	},
}

// GetTaobaoScitemOutercodeGetAPIResponse 从 sync.Pool 获取 TaobaoScitemOutercodeGetAPIResponse
func GetTaobaoScitemOutercodeGetAPIResponse() *TaobaoScitemOutercodeGetAPIResponse {
	return poolTaobaoScitemOutercodeGetAPIResponse.Get().(*TaobaoScitemOutercodeGetAPIResponse)
}

// ReleaseTaobaoScitemOutercodeGetAPIResponse 将 TaobaoScitemOutercodeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoScitemOutercodeGetAPIResponse(v *TaobaoScitemOutercodeGetAPIResponse) {
	v.Reset()
	poolTaobaoScitemOutercodeGetAPIResponse.Put(v)
}
