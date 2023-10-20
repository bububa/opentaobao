package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeVideoUnbindAPIResponse 创意与视频解绑接口 API返回值
// taobao.subway.creative.video.unbind
//
// 将创意与视频解绑
type TaobaoSubwayCreativeVideoUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCreativeVideoUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCreativeVideoUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCreativeVideoUnbindAPIResponseModel).Reset()
}

// TaobaoSubwayCreativeVideoUnbindAPIResponseModel is 创意与视频解绑接口 成功返回结果
type TaobaoSubwayCreativeVideoUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_creative_video_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否解绑成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCreativeVideoUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoSubwayCreativeVideoUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCreativeVideoUnbindAPIResponse)
	},
}

// GetTaobaoSubwayCreativeVideoUnbindAPIResponse 从 sync.Pool 获取 TaobaoSubwayCreativeVideoUnbindAPIResponse
func GetTaobaoSubwayCreativeVideoUnbindAPIResponse() *TaobaoSubwayCreativeVideoUnbindAPIResponse {
	return poolTaobaoSubwayCreativeVideoUnbindAPIResponse.Get().(*TaobaoSubwayCreativeVideoUnbindAPIResponse)
}

// ReleaseTaobaoSubwayCreativeVideoUnbindAPIResponse 将 TaobaoSubwayCreativeVideoUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCreativeVideoUnbindAPIResponse(v *TaobaoSubwayCreativeVideoUnbindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCreativeVideoUnbindAPIResponse.Put(v)
}
