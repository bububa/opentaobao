package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberSyncAPIResponse 会员信息同步 API返回值
// alibaba.member.sync
//
// 会员信息同步
type AlibabaMemberSyncAPIResponse struct {
	model.CommonResponse
	AlibabaMemberSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberSyncAPIResponseModel).Reset()
}

// AlibabaMemberSyncAPIResponseModel is 会员信息同步 成功返回结果
type AlibabaMemberSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMemberSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberSyncAPIResponse)
	},
}

// GetAlibabaMemberSyncAPIResponse 从 sync.Pool 获取 AlibabaMemberSyncAPIResponse
func GetAlibabaMemberSyncAPIResponse() *AlibabaMemberSyncAPIResponse {
	return poolAlibabaMemberSyncAPIResponse.Get().(*AlibabaMemberSyncAPIResponse)
}

// ReleaseAlibabaMemberSyncAPIResponse 将 AlibabaMemberSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberSyncAPIResponse(v *AlibabaMemberSyncAPIResponse) {
	v.Reset()
	poolAlibabaMemberSyncAPIResponse.Put(v)
}
