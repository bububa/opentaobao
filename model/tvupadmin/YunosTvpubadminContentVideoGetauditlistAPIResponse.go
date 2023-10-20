package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentVideoGetauditlistAPIResponse 迎客松视频审核记录查询 API返回值
// yunos.tvpubadmin.content.video.getauditlist
//
// 迎客松视频审核记录查询
type YunosTvpubadminContentVideoGetauditlistAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentVideoGetauditlistAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentVideoGetauditlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentVideoGetauditlistAPIResponseModel).Reset()
}

// YunosTvpubadminContentVideoGetauditlistAPIResponseModel is 迎客松视频审核记录查询 成功返回结果
type YunosTvpubadminContentVideoGetauditlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_video_getauditlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 视频审核列表，json格式
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentVideoGetauditlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentVideoGetauditlistAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentVideoGetauditlistAPIResponse)
	},
}

// GetYunosTvpubadminContentVideoGetauditlistAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentVideoGetauditlistAPIResponse
func GetYunosTvpubadminContentVideoGetauditlistAPIResponse() *YunosTvpubadminContentVideoGetauditlistAPIResponse {
	return poolYunosTvpubadminContentVideoGetauditlistAPIResponse.Get().(*YunosTvpubadminContentVideoGetauditlistAPIResponse)
}

// ReleaseYunosTvpubadminContentVideoGetauditlistAPIResponse 将 YunosTvpubadminContentVideoGetauditlistAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentVideoGetauditlistAPIResponse(v *YunosTvpubadminContentVideoGetauditlistAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentVideoGetauditlistAPIResponse.Put(v)
}
