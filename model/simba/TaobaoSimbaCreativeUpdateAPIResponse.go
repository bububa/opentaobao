package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeUpdateAPIResponse 修改创意 API返回值
// taobao.simba.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
type TaobaoSimbaCreativeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCreativeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCreativeUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaCreativeUpdateAPIResponseModel is 修改创意 成功返回结果
type TaobaoSimbaCreativeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_creative_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创意修改记录对象
	Creativerecord *CreativeRecord `json:"creativerecord,omitempty" xml:"creativerecord,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCreativeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Creativerecord = nil
}

var poolTaobaoSimbaCreativeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCreativeUpdateAPIResponse)
	},
}

// GetTaobaoSimbaCreativeUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaCreativeUpdateAPIResponse
func GetTaobaoSimbaCreativeUpdateAPIResponse() *TaobaoSimbaCreativeUpdateAPIResponse {
	return poolTaobaoSimbaCreativeUpdateAPIResponse.Get().(*TaobaoSimbaCreativeUpdateAPIResponse)
}

// ReleaseTaobaoSimbaCreativeUpdateAPIResponse 将 TaobaoSimbaCreativeUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCreativeUpdateAPIResponse(v *TaobaoSimbaCreativeUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCreativeUpdateAPIResponse.Put(v)
}
