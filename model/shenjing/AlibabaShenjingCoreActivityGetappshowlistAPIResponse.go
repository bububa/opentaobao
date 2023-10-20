package shenjing

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShenjingCoreActivityGetappshowlistAPIResponse 获取神鲸活动列表 API返回值
// alibaba.shenjing.core.activity.getappshowlist
//
// 获取神鲸活动列表
type AlibabaShenjingCoreActivityGetappshowlistAPIResponse struct {
	model.CommonResponse
	AlibabaShenjingCoreActivityGetappshowlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaShenjingCoreActivityGetappshowlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaShenjingCoreActivityGetappshowlistAPIResponseModel).Reset()
}

// AlibabaShenjingCoreActivityGetappshowlistAPIResponseModel is 获取神鲸活动列表 成功返回结果
type AlibabaShenjingCoreActivityGetappshowlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shenjing_core_activity_getappshowlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的结果对象
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaShenjingCoreActivityGetappshowlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaShenjingCoreActivityGetappshowlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaShenjingCoreActivityGetappshowlistAPIResponse)
	},
}

// GetAlibabaShenjingCoreActivityGetappshowlistAPIResponse 从 sync.Pool 获取 AlibabaShenjingCoreActivityGetappshowlistAPIResponse
func GetAlibabaShenjingCoreActivityGetappshowlistAPIResponse() *AlibabaShenjingCoreActivityGetappshowlistAPIResponse {
	return poolAlibabaShenjingCoreActivityGetappshowlistAPIResponse.Get().(*AlibabaShenjingCoreActivityGetappshowlistAPIResponse)
}

// ReleaseAlibabaShenjingCoreActivityGetappshowlistAPIResponse 将 AlibabaShenjingCoreActivityGetappshowlistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaShenjingCoreActivityGetappshowlistAPIResponse(v *AlibabaShenjingCoreActivityGetappshowlistAPIResponse) {
	v.Reset()
	poolAlibabaShenjingCoreActivityGetappshowlistAPIResponse.Put(v)
}
