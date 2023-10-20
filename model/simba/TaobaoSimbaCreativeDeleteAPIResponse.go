package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeDeleteAPIResponse 删除创意 API返回值
// taobao.simba.creative.delete
//
// 删除一个创意
type TaobaoSimbaCreativeDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativeDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCreativeDeleteAPIResponseModel).Reset()
}

// TaobaoSimbaCreativeDeleteAPIResponseModel is 删除创意 成功返回结果
type TaobaoSimbaCreativeDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creative_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除的创意对象
	Creative *Creative `json:"creative,omitempty" xml:"creative,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Creative = nil
}

var poolTaobaoSimbaCreativeDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCreativeDeleteAPIResponse)
	},
}

// GetTaobaoSimbaCreativeDeleteAPIResponse 从 sync.Pool 获取 TaobaoSimbaCreativeDeleteAPIResponse
func GetTaobaoSimbaCreativeDeleteAPIResponse() *TaobaoSimbaCreativeDeleteAPIResponse {
	return poolTaobaoSimbaCreativeDeleteAPIResponse.Get().(*TaobaoSimbaCreativeDeleteAPIResponse)
}

// ReleaseTaobaoSimbaCreativeDeleteAPIResponse 将 TaobaoSimbaCreativeDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCreativeDeleteAPIResponse(v *TaobaoSimbaCreativeDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCreativeDeleteAPIResponse.Put(v)
}
