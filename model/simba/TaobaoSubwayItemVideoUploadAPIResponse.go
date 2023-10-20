package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayItemVideoUploadAPIResponse 创意视频上传 API返回值
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
type TaobaoSubwayItemVideoUploadAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayItemVideoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayItemVideoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayItemVideoUploadAPIResponseModel).Reset()
}

// TaobaoSubwayItemVideoUploadAPIResponseModel is 创意视频上传 成功返回结果
type TaobaoSubwayItemVideoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_item_video_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频信息
	Result *VideoFeedDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayItemVideoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSubwayItemVideoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayItemVideoUploadAPIResponse)
	},
}

// GetTaobaoSubwayItemVideoUploadAPIResponse 从 sync.Pool 获取 TaobaoSubwayItemVideoUploadAPIResponse
func GetTaobaoSubwayItemVideoUploadAPIResponse() *TaobaoSubwayItemVideoUploadAPIResponse {
	return poolTaobaoSubwayItemVideoUploadAPIResponse.Get().(*TaobaoSubwayItemVideoUploadAPIResponse)
}

// ReleaseTaobaoSubwayItemVideoUploadAPIResponse 将 TaobaoSubwayItemVideoUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayItemVideoUploadAPIResponse(v *TaobaoSubwayItemVideoUploadAPIResponse) {
	v.Reset()
	poolTaobaoSubwayItemVideoUploadAPIResponse.Put(v)
}
