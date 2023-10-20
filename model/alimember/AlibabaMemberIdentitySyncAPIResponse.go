package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentitySyncAPIResponse 会员身份信息同步 API返回值
// alibaba.member.identity.sync
//
// 会员身份信息同步
type AlibabaMemberIdentitySyncAPIResponse struct {
	model.CommonResponse
	AlibabaMemberIdentitySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberIdentitySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberIdentitySyncAPIResponseModel).Reset()
}

// AlibabaMemberIdentitySyncAPIResponseModel is 会员身份信息同步 成功返回结果
type AlibabaMemberIdentitySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_identity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaMemberIdentitySyncTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberIdentitySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberIdentitySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentitySyncAPIResponse)
	},
}

// GetAlibabaMemberIdentitySyncAPIResponse 从 sync.Pool 获取 AlibabaMemberIdentitySyncAPIResponse
func GetAlibabaMemberIdentitySyncAPIResponse() *AlibabaMemberIdentitySyncAPIResponse {
	return poolAlibabaMemberIdentitySyncAPIResponse.Get().(*AlibabaMemberIdentitySyncAPIResponse)
}

// ReleaseAlibabaMemberIdentitySyncAPIResponse 将 AlibabaMemberIdentitySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberIdentitySyncAPIResponse(v *AlibabaMemberIdentitySyncAPIResponse) {
	v.Reset()
	poolAlibabaMemberIdentitySyncAPIResponse.Put(v)
}
