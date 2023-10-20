package youkuott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukumediaapivideosnapshotgetAPIResponse 根据视频ID查询视频缩微图 API返回值
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
type YoukumediaapivideosnapshotgetAPIResponse struct {
	model.CommonResponse
	YoukumediaapivideosnapshotgetAPIResponseModel
}

// YoukumediaapivideosnapshotgetAPIResponseModel is 根据视频ID查询视频缩微图 成功返回结果
type YoukumediaapivideosnapshotgetAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_mediaapi_video_snapshot_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据
	ModelList []YoukumediaapivideosnapshotgetStruct `json:"model_list,omitempty" xml:"model_list>youkumediaapivideosnapshotget_struct,omitempty"`
	// 是否成功 true:成功  false:失败（top成功标志）
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
