package cainiaocntec

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse 验证码验证 API返回值
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
type CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse struct {
	model.CommonResponse
	CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponseModel).Reset()
}

// CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponseModel is 验证码验证 成功返回结果
type CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_locallife_communitylife_verifyservicecode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	EMsg string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ECode = ""
	m.Model = ""
	m.EMsg = ""
	m.IsSuccess = false
}

var poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse)
	},
}

// GetCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse 从 sync.Pool 获取 CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse
func GetCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse() *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse {
	return poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse.Get().(*CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse)
}

// ReleaseCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse 将 CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse(v *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse) {
	v.Reset()
	poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse.Put(v)
}
