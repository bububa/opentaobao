package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryVmosCommonCallbackAPIResponse vmos游戏信息采集结果回调通知 API返回值
// alibaba.jym.industry.vmos.common.callback
//
// vmos游戏信息采集结果回调通知
type AlibabaJymIndustryVmosCommonCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaJymIndustryVmosCommonCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymIndustryVmosCommonCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymIndustryVmosCommonCallbackAPIResponseModel).Reset()
}

// AlibabaJymIndustryVmosCommonCallbackAPIResponseModel is vmos游戏信息采集结果回调通知 成功返回结果
type AlibabaJymIndustryVmosCommonCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_vmos_common_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 扩展信息错误
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymIndustryVmosCommonCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
}

var poolAlibabaJymIndustryVmosCommonCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymIndustryVmosCommonCallbackAPIResponse)
	},
}

// GetAlibabaJymIndustryVmosCommonCallbackAPIResponse 从 sync.Pool 获取 AlibabaJymIndustryVmosCommonCallbackAPIResponse
func GetAlibabaJymIndustryVmosCommonCallbackAPIResponse() *AlibabaJymIndustryVmosCommonCallbackAPIResponse {
	return poolAlibabaJymIndustryVmosCommonCallbackAPIResponse.Get().(*AlibabaJymIndustryVmosCommonCallbackAPIResponse)
}

// ReleaseAlibabaJymIndustryVmosCommonCallbackAPIResponse 将 AlibabaJymIndustryVmosCommonCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymIndustryVmosCommonCallbackAPIResponse(v *AlibabaJymIndustryVmosCommonCallbackAPIResponse) {
	v.Reset()
	poolAlibabaJymIndustryVmosCommonCallbackAPIResponse.Put(v)
}
