package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativesGetAPIResponse 批量获得创意 API返回值
// taobao.simba.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
type TaobaoSimbaCreativesGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCreativesGetAPIResponseModel).Reset()
}

// TaobaoSimbaCreativesGetAPIResponseModel is 批量获得创意 成功返回结果
type TaobaoSimbaCreativesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creatives_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创意对象列表
	Creatives []Creative `json:"creatives,omitempty" xml:"creatives>creative,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Creatives = m.Creatives[:0]
}

var poolTaobaoSimbaCreativesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCreativesGetAPIResponse)
	},
}

// GetTaobaoSimbaCreativesGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCreativesGetAPIResponse
func GetTaobaoSimbaCreativesGetAPIResponse() *TaobaoSimbaCreativesGetAPIResponse {
	return poolTaobaoSimbaCreativesGetAPIResponse.Get().(*TaobaoSimbaCreativesGetAPIResponse)
}

// ReleaseTaobaoSimbaCreativesGetAPIResponse 将 TaobaoSimbaCreativesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCreativesGetAPIResponse(v *TaobaoSimbaCreativesGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCreativesGetAPIResponse.Put(v)
}
