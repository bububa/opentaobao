package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse 新房渠道电话数据同步 API返回值
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
type AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectChannelphoneAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectChannelphoneAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectChannelphoneAPIResponseModel is 新房渠道电话数据同步 成功返回结果
type AlibabaAlihouseNewhomeProjectChannelphoneAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_channelphone_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectChannelphoneResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectChannelphoneAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse
func GetAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse() *AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse 将 AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse(v *AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectChannelphoneAPIResponse.Put(v)
}
