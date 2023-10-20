package youkuott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuTvoperatorMediaPageQueryAPIResponse 运营商全量媒资分页查询 API返回值
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
type YoukuTvoperatorMediaPageQueryAPIResponse struct {
	model.CommonResponse
	YoukuTvoperatorMediaPageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuTvoperatorMediaPageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuTvoperatorMediaPageQueryAPIResponseModel).Reset()
}

// YoukuTvoperatorMediaPageQueryAPIResponseModel is 运营商全量媒资分页查询 成功返回结果
type YoukuTvoperatorMediaPageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_tvoperator_media_page_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 数据信息
	Model *YoukuTvoperatorMediaPageQueryModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功 true:成功 false:不成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *YoukuTvoperatorMediaPageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = nil
	m.IsSuccess = false
}

var poolYoukuTvoperatorMediaPageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuTvoperatorMediaPageQueryAPIResponse)
	},
}

// GetYoukuTvoperatorMediaPageQueryAPIResponse 从 sync.Pool 获取 YoukuTvoperatorMediaPageQueryAPIResponse
func GetYoukuTvoperatorMediaPageQueryAPIResponse() *YoukuTvoperatorMediaPageQueryAPIResponse {
	return poolYoukuTvoperatorMediaPageQueryAPIResponse.Get().(*YoukuTvoperatorMediaPageQueryAPIResponse)
}

// ReleaseYoukuTvoperatorMediaPageQueryAPIResponse 将 YoukuTvoperatorMediaPageQueryAPIResponse 保存到 sync.Pool
func ReleaseYoukuTvoperatorMediaPageQueryAPIResponse(v *YoukuTvoperatorMediaPageQueryAPIResponse) {
	v.Reset()
	poolYoukuTvoperatorMediaPageQueryAPIResponse.Put(v)
}
