package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingDiagnosistitleAPIResponse 标题诊断 API返回值
// alibaba.seaking.diagnosistitle
//
// 标题诊断
type AlibabaSeakingDiagnosistitleAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingDiagnosistitleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingDiagnosistitleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingDiagnosistitleAPIResponseModel).Reset()
}

// AlibabaSeakingDiagnosistitleAPIResponseModel is 标题诊断 成功返回结果
type AlibabaSeakingDiagnosistitleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_diagnosistitle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 全大写的单词
	AllUppercaseWords []string `json:"all_uppercase_words,omitempty" xml:"all_uppercase_words>string,omitempty"`
	// 违禁词
	DisableWordList []string `json:"disable_word_list,omitempty" xml:"disable_word_list>string,omitempty"`
	// 重复的词
	DuplicateWordList []string `json:"duplicate_word_list,omitempty" xml:"duplicate_word_list>string,omitempty"`
	// 首字母未大写单词
	NoFirstUppercaseWordList []string `json:"no_first_uppercase_word_list,omitempty" xml:"no_first_uppercase_word_list>string,omitempty"`
	// 拼写错误的单词
	SpellErrorWordList []string `json:"spell_error_word_list,omitempty" xml:"spell_error_word_list>string,omitempty"`
	// 语言质量分(0-5)
	LanguageQualityScore string `json:"language_quality_score,omitempty" xml:"language_quality_score,omitempty"`
	// 总分(0-100)
	TotalScore int64 `json:"total_score,omitempty" xml:"total_score,omitempty"`
	// 是否包含类目主词
	ContainCoreClasses bool `json:"contain_core_classes,omitempty" xml:"contain_core_classes,omitempty"`
	// 标题是否超过长度限制
	OverLengthLimit bool `json:"over_length_limit,omitempty" xml:"over_length_limit,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingDiagnosistitleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AllUppercaseWords = m.AllUppercaseWords[:0]
	m.DisableWordList = m.DisableWordList[:0]
	m.DuplicateWordList = m.DuplicateWordList[:0]
	m.NoFirstUppercaseWordList = m.NoFirstUppercaseWordList[:0]
	m.SpellErrorWordList = m.SpellErrorWordList[:0]
	m.LanguageQualityScore = ""
	m.TotalScore = 0
	m.ContainCoreClasses = false
	m.OverLengthLimit = false
}

var poolAlibabaSeakingDiagnosistitleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingDiagnosistitleAPIResponse)
	},
}

// GetAlibabaSeakingDiagnosistitleAPIResponse 从 sync.Pool 获取 AlibabaSeakingDiagnosistitleAPIResponse
func GetAlibabaSeakingDiagnosistitleAPIResponse() *AlibabaSeakingDiagnosistitleAPIResponse {
	return poolAlibabaSeakingDiagnosistitleAPIResponse.Get().(*AlibabaSeakingDiagnosistitleAPIResponse)
}

// ReleaseAlibabaSeakingDiagnosistitleAPIResponse 将 AlibabaSeakingDiagnosistitleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingDiagnosistitleAPIResponse(v *AlibabaSeakingDiagnosistitleAPIResponse) {
	v.Reset()
	poolAlibabaSeakingDiagnosistitleAPIResponse.Put(v)
}
