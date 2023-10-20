package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowSetexemptauditAPIResponse 迎客松节目设置免审开关 API返回值
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
type YunosTvpubadminContentShowSetexemptauditAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentShowSetexemptauditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowSetexemptauditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentShowSetexemptauditAPIResponseModel).Reset()
}

// YunosTvpubadminContentShowSetexemptauditAPIResponseModel is 迎客松节目设置免审开关 成功返回结果
type YunosTvpubadminContentShowSetexemptauditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_setexemptaudit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设置免审
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentShowSetexemptauditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = false
}

var poolYunosTvpubadminContentShowSetexemptauditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentShowSetexemptauditAPIResponse)
	},
}

// GetYunosTvpubadminContentShowSetexemptauditAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentShowSetexemptauditAPIResponse
func GetYunosTvpubadminContentShowSetexemptauditAPIResponse() *YunosTvpubadminContentShowSetexemptauditAPIResponse {
	return poolYunosTvpubadminContentShowSetexemptauditAPIResponse.Get().(*YunosTvpubadminContentShowSetexemptauditAPIResponse)
}

// ReleaseYunosTvpubadminContentShowSetexemptauditAPIResponse 将 YunosTvpubadminContentShowSetexemptauditAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentShowSetexemptauditAPIResponse(v *YunosTvpubadminContentShowSetexemptauditAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentShowSetexemptauditAPIResponse.Put(v)
}
