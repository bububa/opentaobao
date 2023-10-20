package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativesChangedGetAPIResponse 分页获取修改过的广告创意ID和修改时间 API返回值
// taobao.simba.creatives.changed.get
//
// 分页获取修改过的广告创意ID和修改时间
type TaobaoSimbaCreativesChangedGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativesChangedGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativesChangedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCreativesChangedGetAPIResponseModel).Reset()
}

// TaobaoSimbaCreativesChangedGetAPIResponseModel is 分页获取修改过的广告创意ID和修改时间 成功返回结果
type TaobaoSimbaCreativesChangedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creatives_changed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 广告创意分页对象
	Creatives *CreativePage `json:"creatives,omitempty" xml:"creatives,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativesChangedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Creatives = nil
}

var poolTaobaoSimbaCreativesChangedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCreativesChangedGetAPIResponse)
	},
}

// GetTaobaoSimbaCreativesChangedGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCreativesChangedGetAPIResponse
func GetTaobaoSimbaCreativesChangedGetAPIResponse() *TaobaoSimbaCreativesChangedGetAPIResponse {
	return poolTaobaoSimbaCreativesChangedGetAPIResponse.Get().(*TaobaoSimbaCreativesChangedGetAPIResponse)
}

// ReleaseTaobaoSimbaCreativesChangedGetAPIResponse 将 TaobaoSimbaCreativesChangedGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCreativesChangedGetAPIResponse(v *TaobaoSimbaCreativesChangedGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCreativesChangedGetAPIResponse.Put(v)
}
