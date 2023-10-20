package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest 单个关键词报告 API请求
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 返回详情
	_keywordReportOperation *KeywordReportOperationDto
}

// NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest 初始化AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest {
	return &AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) Reset() {
	r._topContext = nil
	r._keywordReportOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetKeywordReportOperation is KeywordReportOperation Setter
// 返回详情
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
	r._keywordReportOperation = _keywordReportOperation
	r.Set("keyword_report_operation", _keywordReportOperation)
	return nil
}

// GetKeywordReportOperation KeywordReportOperation Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
	return r._keywordReportOperation
}

var poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest()
	},
}

// GetAlibabaScbpAdReportQuerySingleKeywordEffectRequest 从 sync.Pool 获取 AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest
func GetAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest {
	return poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest.Get().(*AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest)
}

// ReleaseAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest 将 AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest(v *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest.Put(v)
}
