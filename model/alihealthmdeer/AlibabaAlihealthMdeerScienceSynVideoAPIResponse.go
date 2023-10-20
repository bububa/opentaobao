package alihealthmdeer

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceSynVideoAPIResponse 视频同步【保存/更新】 API返回值
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
type AlibabaAlihealthMdeerScienceSynVideoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerScienceSynVideoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel).Reset()
}

// AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel is 视频同步【保存/更新】 成功返回结果
type AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_synVideo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用结果code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerScienceSynVideoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = false
}

var poolAlibabaAlihealthMdeerScienceSynVideoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMdeerScienceSynVideoAPIResponse)
	},
}

// GetAlibabaAlihealthMdeerScienceSynVideoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMdeerScienceSynVideoAPIResponse
func GetAlibabaAlihealthMdeerScienceSynVideoAPIResponse() *AlibabaAlihealthMdeerScienceSynVideoAPIResponse {
	return poolAlibabaAlihealthMdeerScienceSynVideoAPIResponse.Get().(*AlibabaAlihealthMdeerScienceSynVideoAPIResponse)
}

// ReleaseAlibabaAlihealthMdeerScienceSynVideoAPIResponse 将 AlibabaAlihealthMdeerScienceSynVideoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMdeerScienceSynVideoAPIResponse(v *AlibabaAlihealthMdeerScienceSynVideoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMdeerScienceSynVideoAPIResponse.Put(v)
}
