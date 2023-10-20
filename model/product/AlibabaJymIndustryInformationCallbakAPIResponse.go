package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryInformationCallbakAPIResponse VMOS回调行业信息系统 API返回值
// alibaba.jym.industry.information.callbak
//
// VMOS回调交易猫行业信息系统
type AlibabaJymIndustryInformationCallbakAPIResponse struct {
	model.CommonResponse
	AlibabaJymIndustryInformationCallbakAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymIndustryInformationCallbakAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymIndustryInformationCallbakAPIResponseModel).Reset()
}

// AlibabaJymIndustryInformationCallbakAPIResponseModel is VMOS回调行业信息系统 成功返回结果
type AlibabaJymIndustryInformationCallbakAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_information_callbak_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 扩展错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymIndustryInformationCallbakAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
}

var poolAlibabaJymIndustryInformationCallbakAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymIndustryInformationCallbakAPIResponse)
	},
}

// GetAlibabaJymIndustryInformationCallbakAPIResponse 从 sync.Pool 获取 AlibabaJymIndustryInformationCallbakAPIResponse
func GetAlibabaJymIndustryInformationCallbakAPIResponse() *AlibabaJymIndustryInformationCallbakAPIResponse {
	return poolAlibabaJymIndustryInformationCallbakAPIResponse.Get().(*AlibabaJymIndustryInformationCallbakAPIResponse)
}

// ReleaseAlibabaJymIndustryInformationCallbakAPIResponse 将 AlibabaJymIndustryInformationCallbakAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymIndustryInformationCallbakAPIResponse(v *AlibabaJymIndustryInformationCallbakAPIResponse) {
	v.Reset()
	poolAlibabaJymIndustryInformationCallbakAPIResponse.Put(v)
}
