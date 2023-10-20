package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotUseroscontrolAPIResponse 物联卡用户管理停开机功能 API返回值
// alibaba.aliqin.fc.iot.useroscontrol
//
// 物联网针对用户级管理停开支持
type AlibabaAliqinFcIotUseroscontrolAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotUseroscontrolAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotUseroscontrolAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotUseroscontrolAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotUseroscontrolAPIResponseModel is 物联卡用户管理停开机功能 成功返回结果
type AlibabaAliqinFcIotUseroscontrolAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_useroscontrol_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotUseroscontrolResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotUseroscontrolAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotUseroscontrolAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotUseroscontrolAPIResponse)
	},
}

// GetAlibabaAliqinFcIotUseroscontrolAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotUseroscontrolAPIResponse
func GetAlibabaAliqinFcIotUseroscontrolAPIResponse() *AlibabaAliqinFcIotUseroscontrolAPIResponse {
	return poolAlibabaAliqinFcIotUseroscontrolAPIResponse.Get().(*AlibabaAliqinFcIotUseroscontrolAPIResponse)
}

// ReleaseAlibabaAliqinFcIotUseroscontrolAPIResponse 将 AlibabaAliqinFcIotUseroscontrolAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotUseroscontrolAPIResponse(v *AlibabaAliqinFcIotUseroscontrolAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotUseroscontrolAPIResponse.Put(v)
}
