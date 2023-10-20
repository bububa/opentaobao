package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse 外部上报游戏属性信息 API返回值
// alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo
//
// 外部上报游戏属性信息
type AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse struct {
	model.CommonResponse
	AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponseModel).Reset()
}

// AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponseModel is 外部上报游戏属性信息 成功返回结果
type AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_industry_outsidegamepropertysync_syncpropertyinfo_response"`
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
	// 上报结果DTO
	Result *OutSideSyncGamePropertyResponseDto `json:"result,omitempty" xml:"result,omitempty"`
	// 接口调用结果: true/false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubCodeType = ""
	m.SubExtraErrMsg = ""
	m.StateCode = ""
	m.ExtraErrMsg = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse)
	},
}

// GetAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse 从 sync.Pool 获取 AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse
func GetAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse() *AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse {
	return poolAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse.Get().(*AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse)
}

// ReleaseAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse 将 AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse(v *AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse) {
	v.Reset()
	poolAlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse.Put(v)
}
