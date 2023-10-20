package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMediaVideoListAPIResponse 获取商家视频列表 API返回值
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
type TaobaoMediaVideoListAPIResponse struct {
	model.CommonResponse
	TaobaoMediaVideoListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMediaVideoListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMediaVideoListAPIResponseModel).Reset()
}

// TaobaoMediaVideoListAPIResponseModel is 获取商家视频列表 成功返回结果
type TaobaoMediaVideoListAPIResponseModel struct {
	XMLName xml.Name `xml:"media_video_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SearchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMediaVideoListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMediaVideoListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMediaVideoListAPIResponse)
	},
}

// GetTaobaoMediaVideoListAPIResponse 从 sync.Pool 获取 TaobaoMediaVideoListAPIResponse
func GetTaobaoMediaVideoListAPIResponse() *TaobaoMediaVideoListAPIResponse {
	return poolTaobaoMediaVideoListAPIResponse.Get().(*TaobaoMediaVideoListAPIResponse)
}

// ReleaseTaobaoMediaVideoListAPIResponse 将 TaobaoMediaVideoListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMediaVideoListAPIResponse(v *TaobaoMediaVideoListAPIResponse) {
	v.Reset()
	poolTaobaoMediaVideoListAPIResponse.Put(v)
}
