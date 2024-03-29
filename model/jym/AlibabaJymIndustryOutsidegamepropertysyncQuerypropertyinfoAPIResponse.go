package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse 外部查询游戏属性库属性信息 API返回值
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
type AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse struct {
	model.CommonResponse
	AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponseModel).Reset()
}

// AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponseModel is 外部查询游戏属性库属性信息 成功返回结果
type AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_outsidegamepropertysync_querypropertyinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子集错误码，提供与业务细节使用
	SubCodeType string `json:"sub_code_type,omitempty" xml:"sub_code_type,omitempty"`
	// 子集错误描述，提供与业务细节使用
	SubExtraErrMsg string `json:"sub_extra_err_msg,omitempty" xml:"sub_extra_err_msg,omitempty"`
	// 错误码枚举
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误详细描述
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 外部查询需要采集属性结构信息响应DTO
	Result *OutSideQueryGamePropertyInfoResponseDto `json:"result,omitempty" xml:"result,omitempty"`
	// 接口调用结果: true/false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubCodeType = ""
	m.SubExtraErrMsg = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse)
	},
}

// GetAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse 从 sync.Pool 获取 AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse
func GetAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse() *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse {
	return poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse.Get().(*AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse)
}

// ReleaseAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse 将 AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse(v *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse) {
	v.Reset()
	poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse.Put(v)
}
