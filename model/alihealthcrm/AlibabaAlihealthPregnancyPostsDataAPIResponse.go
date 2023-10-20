package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyPostsDataAPIResponse 发回帖子信息同步 API返回值
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
type AlibabaAlihealthPregnancyPostsDataAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyPostsDataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyPostsDataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPregnancyPostsDataAPIResponseModel).Reset()
}

// AlibabaAlihealthPregnancyPostsDataAPIResponseModel is 发回帖子信息同步 成功返回结果
type AlibabaAlihealthPregnancyPostsDataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_posts_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyPostsDataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthPregnancyPostsDataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyPostsDataAPIResponse)
	},
}

// GetAlibabaAlihealthPregnancyPostsDataAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPregnancyPostsDataAPIResponse
func GetAlibabaAlihealthPregnancyPostsDataAPIResponse() *AlibabaAlihealthPregnancyPostsDataAPIResponse {
	return poolAlibabaAlihealthPregnancyPostsDataAPIResponse.Get().(*AlibabaAlihealthPregnancyPostsDataAPIResponse)
}

// ReleaseAlibabaAlihealthPregnancyPostsDataAPIResponse 将 AlibabaAlihealthPregnancyPostsDataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPregnancyPostsDataAPIResponse(v *AlibabaAlihealthPregnancyPostsDataAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPregnancyPostsDataAPIResponse.Put(v)
}
