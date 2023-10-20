package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentVideoSubmitauditresultAPIResponse 迎客松提交视频审核结果 API返回值
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
type YunosTvpubadminContentVideoSubmitauditresultAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentVideoSubmitauditresultAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentVideoSubmitauditresultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentVideoSubmitauditresultAPIResponseModel).Reset()
}

// YunosTvpubadminContentVideoSubmitauditresultAPIResponseModel is 迎客松提交视频审核结果 成功返回结果
type YunosTvpubadminContentVideoSubmitauditresultAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_video_submitauditresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentVideoSubmitauditresultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentVideoSubmitauditresultAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentVideoSubmitauditresultAPIResponse)
	},
}

// GetYunosTvpubadminContentVideoSubmitauditresultAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentVideoSubmitauditresultAPIResponse
func GetYunosTvpubadminContentVideoSubmitauditresultAPIResponse() *YunosTvpubadminContentVideoSubmitauditresultAPIResponse {
	return poolYunosTvpubadminContentVideoSubmitauditresultAPIResponse.Get().(*YunosTvpubadminContentVideoSubmitauditresultAPIResponse)
}

// ReleaseYunosTvpubadminContentVideoSubmitauditresultAPIResponse 将 YunosTvpubadminContentVideoSubmitauditresultAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentVideoSubmitauditresultAPIResponse(v *YunosTvpubadminContentVideoSubmitauditresultAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentVideoSubmitauditresultAPIResponse.Put(v)
}
