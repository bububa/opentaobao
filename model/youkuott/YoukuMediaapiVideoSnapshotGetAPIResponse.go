package youkuott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuMediaapiVideoSnapshotGetAPIResponse 根据视频ID查询视频缩微图 API返回值
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
type YoukuMediaapiVideoSnapshotGetAPIResponse struct {
	model.CommonResponse
	YoukuMediaapiVideoSnapshotGetAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuMediaapiVideoSnapshotGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuMediaapiVideoSnapshotGetAPIResponseModel).Reset()
}

// YoukuMediaapiVideoSnapshotGetAPIResponseModel is 根据视频ID查询视频缩微图 成功返回结果
type YoukuMediaapiVideoSnapshotGetAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_mediaapi_video_snapshot_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据
	ModelList []YoukuMediaapiVideoSnapshotGetStruct `json:"model_list,omitempty" xml:"model_list>youku_mediaapi_video_snapshot_get_struct,omitempty"`
	// 是否成功 true:成功  false:失败（top成功标志）
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}

// Reset 清空结构体
func (m *YoukuMediaapiVideoSnapshotGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModelList = m.ModelList[:0]
	m.IsSuccess = ""
	m.MsgInfo = ""
}

var poolYoukuMediaapiVideoSnapshotGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuMediaapiVideoSnapshotGetAPIResponse)
	},
}

// GetYoukuMediaapiVideoSnapshotGetAPIResponse 从 sync.Pool 获取 YoukuMediaapiVideoSnapshotGetAPIResponse
func GetYoukuMediaapiVideoSnapshotGetAPIResponse() *YoukuMediaapiVideoSnapshotGetAPIResponse {
	return poolYoukuMediaapiVideoSnapshotGetAPIResponse.Get().(*YoukuMediaapiVideoSnapshotGetAPIResponse)
}

// ReleaseYoukuMediaapiVideoSnapshotGetAPIResponse 将 YoukuMediaapiVideoSnapshotGetAPIResponse 保存到 sync.Pool
func ReleaseYoukuMediaapiVideoSnapshotGetAPIResponse(v *YoukuMediaapiVideoSnapshotGetAPIResponse) {
	v.Reset()
	poolYoukuMediaapiVideoSnapshotGetAPIResponse.Put(v)
}
