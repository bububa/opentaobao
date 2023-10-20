package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserAuthorizeAPIResponse 用户授权接口 API返回值
// alibaba.idle.isv.user.authorize
//
// 用户授权接口
// 接口调用相关参考文档
// https://www.yuque.com/docs/share/9cd991b7-e3a3-40b6-948c-1835422d0164?# 《闲鱼优品2.0API接入说明》
type AlibabaIdleIsvUserAuthorizeAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvUserAuthorizeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvUserAuthorizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvUserAuthorizeAPIResponseModel).Reset()
}

// AlibabaIdleIsvUserAuthorizeAPIResponseModel is 用户授权接口 成功返回结果
type AlibabaIdleIsvUserAuthorizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_user_authorize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvUserAuthorizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvUserAuthorizeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvUserAuthorizeAPIResponse)
	},
}

// GetAlibabaIdleIsvUserAuthorizeAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvUserAuthorizeAPIResponse
func GetAlibabaIdleIsvUserAuthorizeAPIResponse() *AlibabaIdleIsvUserAuthorizeAPIResponse {
	return poolAlibabaIdleIsvUserAuthorizeAPIResponse.Get().(*AlibabaIdleIsvUserAuthorizeAPIResponse)
}

// ReleaseAlibabaIdleIsvUserAuthorizeAPIResponse 将 AlibabaIdleIsvUserAuthorizeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvUserAuthorizeAPIResponse(v *AlibabaIdleIsvUserAuthorizeAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvUserAuthorizeAPIResponse.Put(v)
}
