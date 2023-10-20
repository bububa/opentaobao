package ott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPlayserviceGetplayurlAPIResponse 获取播放串地址 API返回值
// youku.ott.playservice.getplayurl
//
// 获取播放串地址服务
type YoukuOttPlayserviceGetplayurlAPIResponse struct {
	model.CommonResponse
	YoukuOttPlayserviceGetplayurlAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPlayserviceGetplayurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPlayserviceGetplayurlAPIResponseModel).Reset()
}

// YoukuOttPlayserviceGetplayurlAPIResponseModel is 获取播放串地址 成功返回结果
type YoukuOttPlayserviceGetplayurlAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_playservice_getplayurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PlayUrlV2Vo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPlayserviceGetplayurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYoukuOttPlayserviceGetplayurlAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPlayserviceGetplayurlAPIResponse)
	},
}

// GetYoukuOttPlayserviceGetplayurlAPIResponse 从 sync.Pool 获取 YoukuOttPlayserviceGetplayurlAPIResponse
func GetYoukuOttPlayserviceGetplayurlAPIResponse() *YoukuOttPlayserviceGetplayurlAPIResponse {
	return poolYoukuOttPlayserviceGetplayurlAPIResponse.Get().(*YoukuOttPlayserviceGetplayurlAPIResponse)
}

// ReleaseYoukuOttPlayserviceGetplayurlAPIResponse 将 YoukuOttPlayserviceGetplayurlAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPlayserviceGetplayurlAPIResponse(v *YoukuOttPlayserviceGetplayurlAPIResponse) {
	v.Reset()
	poolYoukuOttPlayserviceGetplayurlAPIResponse.Put(v)
}
